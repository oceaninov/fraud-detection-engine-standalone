package customEcho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/fds22/detection-sys/pkg/customMiddleware"
	"gitlab.com/fds22/detection-sys/pkg/customValidator"
	"go.uber.org/zap"
)

func NewCustomEcho(log *zap.SugaredLogger, vld *customValidator.CustomValidator) *echo.Echo {
	ech := echo.New()
	ech.HideBanner = true
	ech.HidePort = true
	ech.Use(middleware.Recover())
	ech.Use(customMiddleware.RequestLogger(log))
	ech.Validator = vld
	return ech
}
