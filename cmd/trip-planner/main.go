package main

import (
	"trip-planner/internal/logger"
)

func main() {
	l := logger.New(true)
	l.Info("started")
	l.Sync()
	// Create a repository
	//repository := storage.New()

}
