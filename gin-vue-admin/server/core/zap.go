package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger 根据配置构造 zap.Logger
func NewLogger(cfg Zap) *zap.Logger {
	level := zapcore.InfoLevel
	if l, err := zapcore.ParseLevel(cfg.Level); err == nil {
		level = l
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := zapcore.NewJSONEncoder(encoderCfg)
	if cfg.Format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(zapcore.AddSync(newStdout()))), level)
	opts := []zap.Option{zap.AddCaller()}
	if cfg.ShowLine {
		opts = append(opts, zap.AddCallerSkip(0), zap.AddStacktrace(parseStackLevel(cfg.Stacktrace)))
	}
	return zap.New(core, opts...)
}

func parseStackLevel(s string) zapcore.Level {
	l, err := zapcore.ParseLevel(s)
	if err != nil {
		return zapcore.ErrorLevel
	}
	return l
}
