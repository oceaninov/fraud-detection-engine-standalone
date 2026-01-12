package graceful

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Shutdown struct encapsulates all the necessary fields for handling graceful shutdown
type Shutdown struct {
	Port    string             // Port on which the server will run
	Context context.Context    // Context to manage lifecycle of the application
	Timeout time.Duration      // Duration to wait before forcing a shutdown
	Syscall []os.Signal        // OS signals to handle for shutdown
	Logger  *zap.SugaredLogger // Logger for logging events
	Echo    *echo.Echo         // Echo instance for the HTTP server
}

// Run method starts the server and handles graceful shutdown
func (r *Shutdown) Run() {
	const fName = "graceful.Run" // Function name for logging purposes

	// Enable CORS middleware for cross-origin requests
	r.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allow all origins (adjust if needed)
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
	}))

	// Log that the server is starting
	r.Logger.Infow(fName, "reason", "starting server")

	// Goroutine to start the server
	go func(fw *echo.Echo) {
		err := fw.Start(r.Port) // Start the server on the specified port
		if err != nil {
			// Log if the server process is interrupted
			r.Logger.Infow(fName, "reason", "server process was interrupted")
		}
		// Log that the server has started
		r.Logger.Infow(fName, "reason", "server started")
	}(r.Echo)

	wait := make(chan struct{}) // Channel to wait for shutdown to complete

	// Goroutine to handle OS signals and initiate shutdown
	go func() {
		s := make(chan os.Signal, 1)   // Channel to receive OS signals
		signal.Notify(s, r.Syscall...) // Notify on specified signals
		<-s                            // Wait for a signal

		// Log that shutdown is initiated
		r.Logger.Infow(fName, "reason", "shutting down")

		// Timer to force exit after the timeout duration
		timeoutFunc := time.AfterFunc(r.Timeout, func() {
			r.Logger.Warnw(
				fName,
				"reason", "timeout due to waiting for other routines",
				"timeout_elapsed", r.Timeout.Milliseconds(),
			)
			os.Exit(0) // Force exit
		})
		defer timeoutFunc.Stop() // Ensure the timer is stopped if shutdown completes in time

		var wg sync.WaitGroup
		wg.Add(1) // Add a wait group to manage goroutines

		// Goroutine to close the Echo server
		go func() {
			defer wg.Done() // Signal that the goroutine is done
			if err := r.Echo.Close(); err != nil {
				// Log any error during server close
				r.Logger.Errorw(
					fName,
					"reason", err.Error(),
				)
				return
			}
			// Log that the server has shut down gracefully
			r.Logger.Infow(fName, "reason", "process was shutdown gracefully")
		}()

		wg.Wait()   // Wait for all goroutines to complete
		close(wait) // Signal that shutdown is complete
	}()
	<-wait // Block until shutdown is complete
}
