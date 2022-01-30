package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func buildLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return logger
}
