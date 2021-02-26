package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log of type *zap.Logger instantiates the logConfig
	Log *zap.Logger
)

func init() {
	logConf := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if Log, err = logConf.Build(); err != nil {
		panic(err)
	}
}
