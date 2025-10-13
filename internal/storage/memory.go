package storage

import (
	"trip-planner/internal/logger"
	"trip-planner/internal/models"

	"go.uber.org/zap"
)

type Memory struct {
	Users map[string]*models.User
}

func New() *Memory {
	return &Memory{
		Users: map[string]*models.User{},
	}
}

func (r *Memory) AddLocation(userName string, name string, country string) error {
	user, ok := r.Users[userName]
	if !ok {
		logger.Log.Debug("User not in memory\n",
			zap.String("name", userName))
		return ErrUserNotFound
	}

	user.Locations[name] = models.Location{
		Name:    name,
		Country: country,
	}

	return nil
}

func (r *Memory) AddUser(name string) error {
	if _, ok := r.Users[name]; ok {
		return ErrUserExists
	}

	r.Users[name] = &models.User{
		Name:      name,
		Trips:     make([]models.Trip, 0),
		Locations: make(map[string]models.Location),
	}

	return nil
}

func (r *Memory) AddTrip(userName string, tripName string, locationNames []string) error {
	user, ok := r.Users[userName]
	if !ok {
		logger.Log.Debug("User not in memory\n",
			zap.String("name", userName))
		return ErrUserNotFound
	}

	trip := models.Trip{
		Name:      tripName,
		Locations: make([]*models.Location, 0, len(locationNames)),
	}

	for _, locationName := range locationNames {
		location, ok := user.Locations[locationName]
		if !ok {
			logger.Log.Debug("Location not in memory\n",
				zap.String("location", locationName),
				zap.String("trip", tripName))
			return ErrLocationNotFound
		}

		trip.Locations = append(trip.Locations, &location)
	}

	user.Trips = append(user.Trips, trip)

	return nil
}

func (r *Memory) GetUser(userName string) (*models.User, error) {
	if user, ok := r.Users[userName]; ok {
		return user, nil
	}

	return nil, ErrUserNotFound
}
