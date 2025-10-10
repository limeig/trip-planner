package storage

import (
	"fmt"
	"trip-planner/internal/models"
)

type Repository struct {
	Locations map[string]*models.Location
	Users     map[string]*models.User
}

func New() Repository {
	return Repository{Locations: map[string]*models.Location{},
		Users: map[string]*models.User{},
	}
}

func (r *Repository) AddLocation(name string, country string) {
	r.Locations[name] = &models.Location{Name: name, Country: country}
}

func (r *Repository) AddUser(name string) {
	trips := make([]*models.Trip, 0)
	r.Users[name] = &models.User{Name: name, Trips: trips}
}

func (r *Repository) AddTrip(userName string, tripName string, locationNames []string) {
	user, ok := r.Users[userName]
	if !ok {
		fmt.Printf("User %s was not found in memory\n", userName)
		return
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

	user.Trips = append(user.Trips, &trip)
}

func (r *Repository) GetUser(userName string) (*models.User, error) {
	if user, ok := r.Users[userName]; ok {
		return user, nil
	}

	return nil, models.ErrUserNotFound
}
