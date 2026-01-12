package customController

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"gitlab.com/fds22/detection-sys/pkg/errorWrapper"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
)

func bindParams(ectx echo.Context, requestValue interface{}) error {
	// Ensure requestValue is a pointer to a struct
	v := reflect.ValueOf(requestValue)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("requestValue must be a pointer to a struct")
	}

	v = v.Elem() // Dereference the pointer to access the struct
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Handle path parameters
		paramTag := field.Tag.Get("param")
		if paramTag != "" {
			paramValue := ectx.Param(paramTag)
			if paramValue != "" {
				if fieldValue.Kind() == reflect.String {
					fieldValue.SetString(paramValue)
				}
			} else {
				return fmt.Errorf("missing path parameter: %s", paramTag)
			}
		}

		// Handle query parameters
		queryTag := field.Tag.Get("query")
		if queryTag != "" {
			queryValue := ectx.QueryParam(queryTag)
			if queryValue != "" {
				if fieldValue.Kind() == reflect.String {
					fieldValue.SetString(queryValue)
				} else if fieldValue.Kind() == reflect.Int {
					intValue, err := strconv.Atoi(queryValue)
					if err != nil {
						return fmt.Errorf("invalid query parameter: %s must be an integer", queryTag)
					}
					fieldValue.SetInt(int64(intValue))
				} else if fieldValue.Kind() == reflect.Bool {
					boolValue, err := strconv.ParseBool(queryValue)
					if err != nil {
						return fmt.Errorf("invalid query parameter: %s must be a boolean (true/false, 1/0)", queryTag)
					}
					fieldValue.SetBool(boolValue)
				}
			}
		}
	}
	return nil
}

func UploadController(uf interface{}, requestType interface{}) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		const xRequestId = "X-Request-ID"
		const xActionTaker = "X-Action-Taker"
		const xFileBytes = "X-File-Bytes"
		requestId := ectx.Request().Header.Get(xRequestId)
		actionTaker := ectx.FormValue("actionTaker")
		ectx.SetRequest(ectx.Request().WithContext(context.WithValue(ectx.Request().Context(), xRequestId, requestId)))
		ectx.SetRequest(ectx.Request().WithContext(context.WithValue(ectx.Request().Context(), xActionTaker, actionTaker)))

		// Create a new instance of the request type
		requestValue := reflect.New(reflect.TypeOf(requestType)).Interface()

		// Retrieve the uploaded file from the request
		file, err := ectx.FormFile("file")
		if err != nil {
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusBadRequest))
		}

		// Open the file for processing
		src, err := file.Open()
		if err != nil {
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusInternalServerError))
		}
		defer func(src multipart.File) {
			_ = src.Close()
		}(src)

		// Read the file content into a byte slice
		fileContent, err := io.ReadAll(src)
		if err != nil {
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusInternalServerError))
		}

		// Embed the file content into the context
		contextWithAdditionalContent := context.WithValue(ectx.Request().Context(), xFileBytes, fileContent)
		ectx.SetRequest(ectx.Request().WithContext(contextWithAdditionalContent))

		// Prepare to call the use case function with reflection
		ufValue := reflect.ValueOf(uf)
		var in []reflect.Value
		in = append(in, reflect.ValueOf(contextWithAdditionalContent))
		in = append(in, reflect.ValueOf(requestValue))

		// Call the use case function
		out := ufValue.Call(in)

		// Handle the response
		if !out[1].IsNil() {
			var errMsg string
			var reason []string
			var httpCode int
			err = out[1].Interface().(error)
			if ex := errorx.Cast(err); ex != nil {
				// Get error message without cause
				errMsg = ex.Message()
				// Get the underlying cause of the error
				if ex.Cause() != nil {
					reason = append(reason, ex.Cause().Error())
				}
				// Retrieve the HTTP code property
				if hc, ok := errorx.ExtractProperty(ex, errorWrapper.ErrHttpCodeProperty); ok {
					httpCode = hc.(int)
				} else {
					httpCode = http.StatusBadRequest
				}
				errHC, errData := basicObject.NewResponseWithReasonError(errMsg, reason, httpCode)
				return ectx.JSON(errHC, errData)
			}
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusBadRequest))
		}
		return ectx.JSON(http.StatusOK, out[0].Interface())
	}
}

func DownloadController(uf interface{}, requestType interface{}) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		const xRequestId = "X-Request-ID"
		requestId := ectx.Request().Header.Get(xRequestId)
		ectx.SetRequest(ectx.Request().WithContext(context.WithValue(ectx.Request().Context(), xRequestId, requestId)))

		// Create a new instance of the request type
		requestValue := reflect.New(reflect.TypeOf(requestType)).Interface()

		// Binding the request body to request type
		err := ectx.Bind(&requestValue)
		if err != nil {
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusBadRequest))
		}

		// Bind URL path parameters
		err = bindParams(ectx, requestValue)
		if err != nil {
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusBadRequest))
		}

		// Prepare to call the use case function with reflection
		ufValue := reflect.ValueOf(uf)
		var in []reflect.Value
		in = append(in, reflect.ValueOf(ectx.Request().Context()))
		in = append(in, reflect.ValueOf(requestValue))

		// Call the use case function
		out := ufValue.Call(in)

		// Retrieve the file details from the use case function's response
		fileData := out[0].Interface()
		fileContent, ok := fileData.([]byte)
		if !ok {
			httpCode, err := basicObject.NewResponseError("Invalid file data", http.StatusInternalServerError)
			return ectx.JSON(httpCode, err)
		}

		// Set appropriate headers and send the file as a response
		ectx.Response().Header().Set(echo.HeaderContentType, "application/octet-stream")
		ectx.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="downloaded_file"`)
		return ectx.Blob(http.StatusOK, "application/octet-stream", fileContent)
	}
}

func RequestController(uf interface{}, requestType interface{}) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		const xRequestId = "X-Request-ID"
		requestId := ectx.Request().Header.Get(xRequestId)
		ectx.SetRequest(ectx.Request().WithContext(context.WithValue(ectx.Request().Context(), xRequestId, requestId)))

		// Create a new instance of the request type
		requestValue := reflect.New(reflect.TypeOf(requestType)).Interface()

		// Binding the request body to request type
		err := ectx.Bind(&requestValue)
		if err != nil {
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusBadRequest))
		}

		// Bind URL path parameters
		err = bindParams(ectx, requestValue)
		if err != nil {
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusBadRequest))
		}

		// Validate the request body to request type
		const tmplt = "field %s should be %s and its %s"
		errValidate := ectx.Validate(requestValue)
		if errValidate != nil {
			var ve validator.ValidationErrors
			if errors.As(errValidate, &ve) {
				var re []string
				for _, vErrs := range ve {
					fieldName := vErrs.Namespace()
					fieldIsRequired := vErrs.Tag()
					fieldKind := vErrs.Kind()
					re = append(re, fmt.Sprintf(tmplt, fieldName, fieldKind, fieldIsRequired))
				}
				return ectx.JSON(
					basicObject.NewResponseWithReasonError(
						"Bad request", re,
						http.StatusBadRequest,
					),
				)
			}
		}

		// Prepare to call the use case function with reflection
		ufValue := reflect.ValueOf(uf)
		var in []reflect.Value
		in = append(in, reflect.ValueOf(ectx.Request().Context()))
		in = append(in, reflect.ValueOf(requestValue))

		// Call the use case function
		out := ufValue.Call(in)

		// Handle the response
		if !out[1].IsNil() {
			var errMsg string
			var reason []string
			var httpCode int
			err = out[1].Interface().(error)
			if ex := errorx.Cast(err); ex != nil {
				// Get error message without cause
				errMsg = ex.Message()
				// Get the underlying cause of the error
				if ex.Cause() != nil {
					reason = append(reason, ex.Cause().Error())
				}
				// Retrieve the HTTP code property
				if hc, ok := errorx.ExtractProperty(ex, errorWrapper.ErrHttpCodeProperty); ok {
					httpCode = hc.(int)
				} else {
					httpCode = http.StatusBadRequest
				}
				errHC, errData := basicObject.NewResponseWithReasonError(errMsg, reason, httpCode)
				return ectx.JSON(errHC, errData)
			}
			return ectx.JSON(basicObject.NewResponseError(err.Error(), http.StatusBadRequest))
		}
		return ectx.JSON(http.StatusOK, out[0].Interface())
	}
}
