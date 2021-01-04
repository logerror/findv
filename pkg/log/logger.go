package log

import (
	"e.welights.net/devsecops/findv/pkg/utils/dateutils"
	"time"

	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOptions struct {
	Level            string
	OutputPaths      string
	ErrorOutputPaths string
	SentryDsn        string
}

var levelMapping = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

func InitLogger(level string) *zap.Logger {
	return InitLoggerWithOptions(LoggerOptions{Level: level})
}

// InitDefaultLogger creates a default logger with level error
func InitDefaultLogger() *zap.Logger {
	return InitLoggerWithOptions(LoggerOptions{})
}

func InitLoggerWithOptions(options LoggerOptions) *zap.Logger {
	config := initConfig(options)

	var err error
	zapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic("Init logger failed:" + err.Error())
	}
	logger = zapLogger
	zap.ReplaceGlobals(logger)

	return logger
}

func Sync() {
	if logger != nil {
		_ = logger.Sync()
	}
}

func modifyToSentryLogger(log *zap.Logger, dsn string) *zap.Logger {
	cfg := zapsentry.Configuration{
		Level: zapcore.ErrorLevel,
	}

	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(dsn))
	if err != nil {
		log.Error("failed to init zap", zap.Error(err))
	}
	return zapsentry.AttachCoreToLogger(core, log)
}

func initConfig(options LoggerOptions) zap.Config {
	config := zap.NewProductionConfig()
	// 默认 error
	if options.Level == "" {
		options.Level = "error"
	}

	config.Level = getLevel(options.Level)

	std := []string{"stdout"}
	err := []string{"stderr"}

	if options.OutputPaths != "" {
		std = append(std, options.OutputPaths)
	}

	if options.OutputPaths != "" {
		err = append(err, options.ErrorOutputPaths)
	}

	config.OutputPaths = std
	config.ErrorOutputPaths = err

	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.NameKey = "name"
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		encodeTimeLayout(t, dateutils.DateTimestampLayout, enc)
	}

	logLevel = config.Level.String()

	return config
}

func getLevel(level string) zap.AtomicLevel {
	zapLevel := levelMapping[level]
	return zap.NewAtomicLevelAt(zapLevel)
}

func encodeTimeLayout(t time.Time, layout string, enc zapcore.PrimitiveArrayEncoder) {
	type appendTimeEncoder interface {
		AppendTimeLayout(time.Time, string)
	}

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, layout)
		return
	}

	enc.AppendString(t.Format(layout))
}
