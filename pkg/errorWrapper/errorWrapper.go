package errorWrapper

import (
	"errors"
	"fmt"
	"github.com/joomcode/errorx"
	"gitlab.com/fds22/detection-sys/pkg/basicObject"
	"net/http"
)

var (
	ErrNamespace = errorx.NewNamespace("fdsys")
	ErrBase      = errorx.NewType(ErrNamespace, "base")
)

var (
	ErrCodeProperty     = errorx.RegisterProperty("code")
	ErrHttpCodeProperty = errorx.RegisterProperty("httpCode")
)

var (
	ErrRepository    = ErrBase.New("datasource").WithProperty(ErrHttpCodeProperty, http.StatusInternalServerError)
	ErrExternalAPI   = ErrBase.New("external").WithProperty(ErrHttpCodeProperty, http.StatusBadGateway)
	ErrBusinessLogic = ErrBase.New("business").WithProperty(ErrHttpCodeProperty, http.StatusBadGateway)
)

type wrapper struct {
	code string
}

type Wrapper interface {
	WrapRepositoryError(err error) error
	WrapExternalError(err error) error
	WrapBusinessError(err error) error
	WrapErrorFromResponse(obj basicObject.ResponseError) error
	FCode() string
}

func NewWrapper(code string) Wrapper {
	return &wrapper{code: code}
}

func (w *wrapper) FCode() string {
	return w.code
}

func (w *wrapper) WrapRepositoryError(err error) error {
	errSegment := fmt.Sprintf("%s %s", ErrRepository.Error(), "operation failed")
	errCauses := errors.New(fmt.Sprintf("[%s]", err.Error()))
	return errorx.Decorate(errCauses, errSegment).WithProperty(ErrCodeProperty, w.code)
}

func (w *wrapper) WrapExternalError(err error) error {
	errSegment := fmt.Sprintf("%s %s", ErrExternalAPI.Error(), "operation failed")
	errCauses := errors.New(fmt.Sprintf("[%s]", err.Error()))
	return errorx.Decorate(errCauses, errSegment).WithProperty(ErrCodeProperty, w.code)
}

func (w *wrapper) WrapBusinessError(err error) error {
	errSegment := fmt.Sprintf("%s %s", ErrBusinessLogic.Error(), "operation violation")
	errCauses := errors.New(fmt.Sprintf("[%s]", err.Error()))
	return errorx.Decorate(errCauses, errSegment).WithProperty(ErrCodeProperty, w.code)
}

func (w *wrapper) WrapErrorFromResponse(obj basicObject.ResponseError) error {
	errSegment := fmt.Sprintf("%s", obj.Message)
	errCauses := errors.New(fmt.Sprintf("%s", errors.New(obj.Message)))
	return errorx.Decorate(errCauses, errSegment).WithProperty(ErrCodeProperty, w.code)
}
