package main

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/fds22/detection-sys/pkg/environments"
	"gitlab.com/fds22/detection-sys/pkg/graceful"
	"gitlab.com/fds22/detection-sys/src/di"
	"go.uber.org/zap"
)

func main() {
	if err := di.Container.Invoke(func(
		log *zap.SugaredLogger,
		ech *echo.Echo,
		envs *environments.Envs,
	) {
		const fName = "entry.main.MainFunction"
		log.Infow(fName, "reason", "execution started")

		log.Infow(fName, "reason", "Base url proxy sso", "url", envs.ProxyAdapterHost)
		log.Infow(fName, "reason", "Redirection uri proxy sso", "uri", envs.ApplicationAzureSSORedirectUri)

		// starting service
		ctx := context.Background()
		defer ctx.Done()

		// starting server through graceful protocol
		gfs := graceful.Shutdown{
			Context: ctx,
			Port:    fmt.Sprintf(":%s", envs.ApplicationPort),
			Echo:    ech,
			Logger:  log,
			Timeout: 15 * time.Second,
			Syscall: []os.Signal{
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGHUP,
			},
		}
		gfs.Run()
	}); err != nil {
		panic(err)
	}
}
