package main

import (
	"trip-planner/internal/logger"
)

func main() {
	l := logger.New(true)
	l.Info("started")

	// Create a repository
	//repository := storage.New()

}
