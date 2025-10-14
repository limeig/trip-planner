package logger

import "go.uber.org/zap"

func New(devMode bool) *zap.Logger {
	var log *zap.Logger
	if devMode {
		log, _ = zap.NewDevelopment()
	} else {
		log, _ = zap.NewProduction()
	}

	return log
}
