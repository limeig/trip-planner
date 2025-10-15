package storage

import (
	"testing"
	"trip-planner/internal/logger"
	"trip-planner/internal/storage/errors"
)

// TestAddUser tests adding a new user to the memory storage.
// Positive scenarios
func TestAddUser(t *testing.T) {
	repo := Init(logger.New(true))

	userName := "Alina"
	err := repo.AddUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	user, err := repo.GetUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	if user.Name != userName {
		t.Fatalf("expected user name to be %s, got %s", userName, user.Name)
	}
}

// Negative scenarios
func TestAddDuplicateUser(t *testing.T) {
	repo := Init(logger.New(true))

	err := repo.AddUser("Bob")
	if err != nil {
		t.Fatal(err)
	}

	err = repo.AddUser("Bob")
	if err != errors.ErrUserExists {
		t.Fatalf("expected %s, got %v", errors.ErrUserExists, err)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	repo := Init(logger.New(true))
	_, err := repo.GetUser("Charlie")
	if err != errors.ErrUserNotFound {
		t.Fatalf("expected %s, got %v", errors.ErrUserNotFound, err)
	}
}

// Test AddLocation
// Positive scenarios
func TestAddLocation(t *testing.T) {
	repo := Init(logger.New(true))

	userName := "Diana"
	err := repo.AddUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.AddLocation(userName, "Paris", "France")
	if err != nil {
		t.Fatal(err)
	}

	user, err := repo.GetUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	location, ok := user.Locations["Paris"]
	if !ok {
		t.Fatalf("expected location Paris to be added, but it was not found")
	}

	if location.Country != "France" {
		t.Fatalf("expected location country to be France, got %s", location.Country)
	}
}

// Negative scenarios
func TestAddLocationToNonExistentUser(t *testing.T) {
	repo := Init(logger.New(true))

	err := repo.AddLocation("NonExistentUser", "Paris", "France")
	if err != errors.ErrUserNotFound {
		t.Fatalf("expected %s, got %v", errors.ErrUserNotFound, err)
	}
}

// Test AddTrip
// Positive scenarios
func TestAddTrip(t *testing.T) {
	repo := Init(logger.New(true))

	userName := "Diana"
	err := repo.AddUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.AddLocation(userName, "London", "UK")
	if err != nil {
		t.Fatal(err)
	}

	err = repo.AddLocation(userName, "Edinburgh", "UK")
	if err != nil {
		t.Fatal(err)
	}

	err = repo.AddTrip(userName, "UK Trip", []string{"London", "Edinburgh"})
	if err != nil {
		t.Fatal(err)
	}

	user, err := repo.GetUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	if len(user.Trips) != 1 {
		t.Fatalf("expected 1 trip, got %d", len(user.Trips))
	}

	if user.Trips[0].Name != "UK Trip" {
		t.Fatalf("expected trip name to be UK Trip, got %s", user.Trips[0].Name)
	}

	if len(user.Trips[0].Locations) != 2 {
		t.Fatalf("expected 2 locations in the trip, got %d", len(user.Trips[0].Locations))
	}

	if user.Trips[0].Locations[0].Name != "London" || user.Trips[0].Locations[1].Name != "Edinburgh" {
		t.Fatalf("expected trip locations to be London and Edinburgh, got %v", user.Trips[0].Locations)
	}
}

// Negative scenarios
func TestAddTripToNonExistentUser(t *testing.T) {
	repo := Init(logger.New(true))

	err := repo.AddTrip("NonExistentUser", "Europe Trip", []string{"Paris", "Berlin"})
	if err != errors.ErrUserNotFound {
		t.Fatalf("expected %s, got %v", errors.ErrUserNotFound, err)
	}
}

func TestAddTripWithNonExistentLocation(t *testing.T) {
	repo := Init(logger.New(true))

	userName := "Diana"
	err := repo.AddUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.AddLocation(userName, "London", "UK")
	if err != nil {
		t.Fatal(err)
	}

	err = repo.AddTrip(userName, "UK Trip", []string{"London", "NonExistentCity"})

	if err != errors.ErrLocationNotFound {
		t.Fatalf("expected %s, got %v", errors.ErrLocationNotFound, err)
	}

	user, err := repo.GetUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	if len(user.Trips) != 0 {
		t.Fatalf("expected 0 trips, got %d", len(user.Trips))
	}
}
