package logger

import (
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// 初始化日志记录器
func init() {
	logDir := path.Join("cmd", "logs")
	_ = os.MkdirAll(logDir, 0777)
	logPath := path.Join(logDir, "log.log")

	fileWriter, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(fileWriter),
		zapcore.DebugLevel,
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// 记录info级别日志
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// 记录debug级别日志
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}
