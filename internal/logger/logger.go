package logger

import "go.uber.org/zap"

var Log *zap.Logger

func Init(devMode bool) {
	if devMode {
		Log, _ = zap.NewDevelopment()
	} else {
		Log, _ = zap.NewProduction()
	}
}
