package storage

import (
	"fmt"
	"trip-planner/internal/models"
)

type Memory struct {
	Locations map[string]*models.Location
	Users     map[string]*models.User
}

func New() Memory {
	return Memory{Locations: map[string]*models.Location{},
		Users: map[string]*models.User{},
	}
}

func (r *Memory) AddLocation(name string, country string) error {
	r.Locations[name] = &models.Location{Name: name, Country: country}
	return nil
}

func (r *Memory) AddUser(name string) error {
	trips := make([]models.Trip, 0)
	r.Users[name] = &models.User{Name: name, Trips: trips}
	return nil
}

func (r *Memory) AddTrip(userName string, tripName string, locationNames []string) error {
	user, ok := r.Users[userName]
	if !ok {
		fmt.Printf("User %s was not found in memory\n", userName)
		return models.ErrUserNotFound
	}

	var trip models.Trip
	for _, locationName := range locationNames {
		location, ok := r.Locations[locationName]
		if !ok {
			fmt.Printf("Location %s for the trip %s was not found in memory\n", locationName, tripName)
			continue
		}

		trip.Locations = append(trip.Locations, location)
	}

	user.Trips = append(user.Trips, trip)

	return nil
}

func (r *Memory) GetUser(userName string) (*models.User, error) {
	if user, ok := r.Users[userName]; ok {
		return user, nil
	}

	return nil, models.ErrUserNotFound
}
