package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var logLevel string

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}

func With(fields ...zapcore.Field) *zap.Logger {
	return logger.With(fields...)
}

func Named(s string) *zap.Logger {
	return logger.Named(s)
}

func Level() string {
	return logLevel
}

func IsDebug() bool {
	return Level() == zap.DebugLevel.String()
}
