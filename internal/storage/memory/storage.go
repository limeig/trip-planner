package memory

import (
	"trip-planner/internal/models"
	"trip-planner/internal/storage/errors"

	"go.uber.org/zap"
)

type Memory struct {
	Users map[string]*models.User
	log   *zap.Logger
}

// New creates a new in-memory storage
func New(log *zap.Logger) *Memory {
	return &Memory{
		Users: map[string]*models.User{},
		log:   log.Named("memory-storage"),
	}
}

// AddLocation adds a new Location to a User's list Locations
// Returns an error if the user does not exist
func (r *Memory) AddLocation(userName string, name string, country string) error {
	user, ok := r.Users[userName]
	if !ok {
		r.log.Debug("User not in memory\n",
			zap.String("name", userName))
		return errors.ErrUserNotFound
	}

	user.Locations[name] = models.Location{
		Name:    name,
		Country: country,
	}

	return nil
}

// AddUser adds a new User to the memory storage
// Returns an error if the user already exists
func (r *Memory) AddUser(name string) error {
	if _, ok := r.Users[name]; ok {
		return errors.ErrUserExists
	}

	r.Users[name] = &models.User{
		Name:      name,
		Trips:     make([]models.Trip, 0),
		Locations: make(map[string]models.Location),
	}

	return nil
}

// AddTrip adds a new Trip to a User's list of Trips
// Returns an error if the user does not exist or if any of the locations do not exist
func (r *Memory) AddTrip(userName string, tripName string, locationNames []string) error {
	user, ok := r.Users[userName]
	if !ok {
		r.log.Debug("User not in memory\n",
			zap.String("name", userName))
		return errors.ErrUserNotFound
	}

	trip := models.Trip{
		Name:      tripName,
		Locations: make([]*models.Location, 0, len(locationNames)),
	}

	for _, locationName := range locationNames {
		location, ok := user.Locations[locationName]
		if !ok {
			r.log.Debug("Location not in memory\n",
				zap.String("location", locationName),
				zap.String("trip", tripName))
			return errors.ErrLocationNotFound
		}

		trip.Locations = append(trip.Locations, &location)
	}

	user.Trips = append(user.Trips, trip)

	return nil
}

// GetUser retrieves a User by name
// Returns an error if the user does not exist
func (r *Memory) GetUser(userName string) (*models.User, error) {
	if user, ok := r.Users[userName]; ok {
		return user, nil
	}

	return nil, errors.ErrUserNotFound
}
