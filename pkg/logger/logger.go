package logger

import (
	"github.com/natefinch/lumberjack"
	"gitlab.com/fds22/detection-sys/pkg/environments"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func NewZapLogger(e *environments.Envs) *zap.SugaredLogger {
	// encode config mode checking and setup
	var encoderCfg zapcore.EncoderConfig
	if strings.Contains(e.SetMode, "dev") {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	}
	if strings.Contains(e.SetMode, "stg") ||
		strings.Contains(e.SetMode, "prd") {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// zap logging config setup using previous encoder config
	var config zap.Config
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.Development = false
	config.DisableCaller = false
	config.DisableStacktrace = false
	config.Sampling = nil
	config.EncoderConfig = encoderCfg
	config.Encoding = e.LogFormatter
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.InitialFields = map[string]interface{}{
		"service_pid":     os.Getpid(),
		"service_name":    e.ApplicationName,
		"service_version": e.ApplicationVersion,
	}

	// Set up file output using Lumberjack
	fileSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   e.LogFilePath,  // log file path
		MaxSize:    e.LogMaxSize,   // maximum megabytes
		MaxBackups: e.LogMaxBackup, // maximum backup
		MaxAge:     e.LogMaxAge,    // maximum days
		Compress:   e.LogCompress,  // compress old logs
	})

	// Create file core
	fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), fileSyncer, config.Level)
	core := zapcore.NewTee(
		zap.Must(config.Build()).Core(),
		fileCore.With([]zap.Field{
			zap.Int("service_pid", os.Getpid()),
			zap.String("service_name", e.ApplicationName),
			zap.String("service_version", e.ApplicationVersion),
		}),
	)

	// setup sugared log
	sugared := zap.New(core).Sugar()
	defer func(sugar *zap.SugaredLogger) {
		_ = sugar.Sync()
	}(sugared)
	return sugared
}
