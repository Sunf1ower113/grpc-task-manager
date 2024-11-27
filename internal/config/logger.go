package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initializes and returns a Zap logger based on the provided configuration.
// It sets the log level, output paths, and log format.
func InitLogger(cfg *LoggerConfig) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	zapCfg := zap.Config{
		Level:            lvl,
		Development:      false,
		Encoding:         "json",
		OutputPaths:      cfg.OutputPaths,
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			MessageKey:     "message",
			LevelKey:       "level",
			CallerKey:      "caller",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
	}

	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
