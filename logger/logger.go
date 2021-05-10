package logger

import (
	"github.com/aditya43/golang-bookstore_users-api/utils/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	if env.Get("RELEASE_ENVIRONMENT") == "release" {
		buildProductionLogger()
	}

	if env.Get("RELEASE_ENVIRONMENT") != "release" {
		buildDevelopmentLogger()
	}
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	_ = log.Sync()
}

func Error(err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(err.Error(), tags...)
	_ = log.Sync()
}

func buildDevelopmentLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	opts := []zap.Option{
		zap.AddCallerSkip(1),
		zap.AddCaller(),
	}

	var err error
	if log, err = config.Build(opts...); err != nil {
		panic(err)
	}
}

func buildProductionLogger() {
	config := zap.NewProductionConfig()
	config.Encoding = "json"
	opts := []zap.Option{
		zap.AddCallerSkip(1),
		zap.AddCaller(),
	}

	var err error
	if log, err = config.Build(opts...); err != nil {
		panic(err)
	}

	// // Custom Config
	// logConfig := zap.Config{
	// 	OutputPaths: []string{"stdout"},
	// 	Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
	// 	Encoding:    "json",
	// 	DisableCaller: false,
	// 	EncoderConfig: zapcore.EncoderConfig{
	// 		LevelKey:    "level",
	// 		TimeKey:     "time",
	// 		MessageKey:  "msg",
	// 		EncodeTime:  zapcore.ISO8601TimeEncoder,
	// 		EncodeLevel:  zapcore.LowercaseLevelEncoder,
	// 		EncodeCaller: zapcore.ShortCallerEncoder,
	// 	},
	// }

	// var err error
	// if log, err = logConfig.Build(); err != nil {
	// 	panic(err)
	// }
}
