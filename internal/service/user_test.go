package service

import (
	"errors"
	"testing"
	"trip-planner/internal/logger"
	"trip-planner/internal/storage"
)

// + Positive scenarios
func TestCreateUser(t *testing.T) {
	log := logger.New(true)
	service := New(storage.Init(log), log)

	userName := "Alice"
	err := service.CreateUser(userName)
	if err != nil {
		t.Fatal(err)
	}
}

// - Negative scenarios
func TestCreateUserEmptyName(t *testing.T) {
	log := logger.New(true)
	service := New(storage.Init(log), log)

	err := service.CreateUser("")
	if !errors.Is(err, ErrUserNameEmpty) {
		t.Fatalf("expected %s, got %v", ErrUserNameEmpty, err)
	}
}
