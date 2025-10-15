package storage

import (
	"trip-planner/internal/models"
	"trip-planner/internal/storage/memory"

	"go.uber.org/zap"
)

type Storage interface {
	AddLocation(string, string, string) error
	AddUser(string) error
	AddTrip(string, string, []string) error

	GetUser(string) (*models.User, error)
}

func Init(log *zap.Logger) Storage {
	// Here you can switch between different storage implementations
	// For example, you can return a memory storage or a database storage
	// return memory.New()
	return memory.New(log)
}
