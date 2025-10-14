package storage

import (
	"trip-planner/internal/models"
)

type Storage interface {
	AddLocation(string, string, string) error
	AddUser(string) error
	AddTrip(string, string, []string) error

	GetUser(string) (*models.User, error)
}
