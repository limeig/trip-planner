package storage

import (
	"fmt"
	"trip-planner/internal/models"
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
		fmt.Printf("User %s was not found in memory\n", userName)
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
		fmt.Printf("User %s was not found in memory\n", userName)
		return ErrUserNotFound
	}

	var trip models.Trip
	for _, locationName := range locationNames {
		location, ok := user.Locations[locationName]
		if !ok {
			fmt.Printf("Location %s for the trip %s was not found in memory\n", locationName, tripName)
			continue
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
