package customMiddleware

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gitlab.com/fds22/detection-sys/pkg/defaultHeaders"
	"gitlab.com/fds22/detection-sys/pkg/utility"
	"go.uber.org/zap"
	"io"
	"net/http"
)

// extractReadsHttpOperation extract http read operations data and content
func extractReadsHttpOperation(next echo.HandlerFunc, ectx echo.Context, log *zap.SugaredLogger) error {
	const fName = "middleware.RequestLogger"
	requestId := ectx.Request().Header.Get(defaultHeaders.XRequestId)
	logging := log.With("request_id", requestId)

	// Manually read response body using custom response writer
	responseBody := new(bytes.Buffer)
	multiWriter := io.MultiWriter(ectx.Response().Writer, responseBody)
	writer := &utility.RPWriter{ResponseWriter: ectx.Response().Writer, Writer: multiWriter}
	ectx.Response().Writer = writer

	// Extract path parameters
	pathParams := make(map[string]interface{})
	for _, param := range ectx.ParamNames() {
		pathParams[param] = ectx.Param(param)
	}

	// Extract query parameters
	queryParams := make(map[string]interface{})
	for key, values := range ectx.QueryParams() {
		if len(values) > 0 {
			queryParams[key] = values[0] // Get the first value if there are multiple
		}
	}

	// Combine path and query parameters
	combinedParams := make(map[string]interface{})
	for k, v := range pathParams {
		combinedParams[k] = v
	}
	for k, v := range queryParams {
		combinedParams[k] = v
	}

	// Unmarshal request body bytes into a map of string interfaces
	bodyBytes, _ := json.Marshal(combinedParams)
	var mapRequest map[string]interface{}
	body := bytes.NewBuffer(bodyBytes).String()
	byteReq := []byte(body)
	_ = json.Unmarshal(byteReq, &mapRequest)
	logging.Infow(fName, "request_body", mapRequest)

	// Call the next handler in the chain
	err := next(ectx)

	// Unmarshal response body bytes into a map of string interfaces
	var mapResponse map[string]interface{}
	byteRes := []byte(responseBody.String())
	_ = json.Unmarshal(byteRes, &mapResponse)
	logging.Infow(fName, "response_data", mapResponse)

	// returning result
	return err
}

// extractWritesHttpOperation extract http write operations data and content
func extractWritesHttpOperation(next echo.HandlerFunc, ectx echo.Context, log *zap.SugaredLogger) error {
	const fName = "middleware.RequestLogger"
	requestId := ectx.Request().Header.Get(defaultHeaders.XRequestId)
	logging := log.With("request_id", requestId)

	// Manually read request body
	var bodyBytes []byte
	if ectx.Request().Body != nil {
		bodyBytes, _ = io.ReadAll(ectx.Request().Body)
	}

	// Manually read response body using custom response writer
	responseBody := new(bytes.Buffer)
	multiWriter := io.MultiWriter(ectx.Response().Writer, responseBody)
	writer := &utility.RPWriter{ResponseWriter: ectx.Response().Writer, Writer: multiWriter}
	ectx.Response().Writer = writer

	// Rewriting request content for the next middleware
	ectx.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Unmarshal request body bytes into a map of string interfaces
	var mapRequest map[string]interface{}
	body := bytes.NewBuffer(bodyBytes).String()
	byteReq := []byte(body)
	_ = json.Unmarshal(byteReq, &mapRequest)
	logging.Infow(fName, "request_body", mapRequest)

	// Call the next handler in the chain
	err := next(ectx)

	// Unmarshal response body bytes into a map of string interfaces
	var mapResponse map[string]interface{}
	byteRes := []byte(responseBody.String())
	_ = json.Unmarshal(byteRes, &mapResponse)
	logging.Infow(fName, "response_data", mapResponse)

	// returning result
	return err
}

// RequestLogger middleware logs incoming requests and outgoing responses
func RequestLogger(log *zap.SugaredLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			switch {
			case ectx.Request().Method == http.MethodGet ||
				ectx.Request().Method == http.MethodDelete ||
				ectx.Request().Method == http.MethodHead:
				return extractReadsHttpOperation(next, ectx, log)
			case ectx.Request().Method == http.MethodPost ||
				ectx.Request().Method == http.MethodPut ||
				ectx.Request().Method == http.MethodPatch ||
				ectx.Request().Method == http.MethodDelete ||
				ectx.Request().Method == http.MethodHead:
				return extractWritesHttpOperation(next, ectx, log)
			default:
				// Call the next handler in the chain
				err := next(ectx)
				return err
			}
		}
	}
}
