package main

import (
	"trip-planner/internal/logger"
)

func main() {
	logger.Init(true)
	logger.Log.Info("Trip Planner started")

	// Create a repository
	//repository := storage.New()

}
