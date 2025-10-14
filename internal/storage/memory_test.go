package storage

import (
	"testing"
	"trip-planner/internal/logger"
)

// TestAddUser tests adding a new user to the memory storage.
// Positive scenarios
func TestAddUser(t *testing.T) {
	log := logger.New(true)
	repo := New(log)

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
	log := logger.New(true)
	repo := New(log)
	err := repo.AddUser("Bob")
	if err != nil {
		t.Fatal(err)
	}

	err = repo.AddUser("Bob")
	if err != ErrUserExists {
		t.Fatalf("expected %s, got %v", ErrUserExists, err)
	}
}

func TestGetNonExistentUser(t *testing.T) {
	log := logger.New(true)
	repo := New(log)
	_, err := repo.GetUser("Charlie")
	if err != ErrUserNotFound {
		t.Fatalf("expected %s, got %v", ErrUserNotFound, err)
	}
}

// Test AddLocation
// Positive scenarios
func TestAddLocation(t *testing.T) {
	log := logger.New(true)
	repo := New(log)
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
	log := logger.New(true)
	repo := New(log)
	err := repo.AddLocation("NonExistentUser", "Paris", "France")
	if err != ErrUserNotFound {
		t.Fatalf("expected %s, got %v", ErrUserNotFound, err)
	}
}

// Test AddTrip
// Positive scenarios
func TestAddTrip(t *testing.T) {
	log := logger.New(true)
	repo := New(log)
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
	log := logger.New(true)
	repo := New(log)
	err := repo.AddTrip("NonExistentUser", "Europe Trip", []string{"Paris", "Berlin"})
	if err != ErrUserNotFound {
		t.Fatalf("expected %s, got %v", ErrUserNotFound, err)
	}
}

func TestAddTripWithNonExistentLocation(t *testing.T) {
	log := logger.New(true)
	repo := New(log)
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

	if err != ErrLocationNotFound {
		t.Fatalf("expected %s, got %v", ErrLocationNotFound, err)
	}

	user, err := repo.GetUser(userName)
	if err != nil {
		t.Fatal(err)
	}

	if len(user.Trips) != 0 {
		t.Fatalf("expected 0 trips, got %d", len(user.Trips))
	}
}
