package service

import (
	"trip-planner/internal/models"
	"trip-planner/internal/storage"

	"go.uber.org/zap"
)

type UserService struct {
	Storage storage.Storage
	Logger  *zap.Logger
}

func New(storage storage.Storage, log zap.Logger) *UserService {
	return &UserService{
		Storage: storage,
		Logger:  log.Named("user-service"),
	}
}

// CreateUser adds a new user to the repository and logs the operation
// Returns an error if the user already exists or if the name is empty
func (s *UserService) CreateUser(name string) error {
	if name == "" {
		s.Logger.Error("User name is empty")
		return ErrUserNameEmpty
	}

	if err := s.Storage.AddUser(name); err != nil {
		s.Logger.Error("Failed to create user", zap.String("name", name), zap.Error(err))
		return err
	}

	s.Logger.Debug("Created user", zap.String("name", name))
	return nil
}

// GetUser retrieves a user by name and logs the operation
// Returns an error if the user does not exist or if the name is empty
func (s *UserService) GetUser(name string) (*models.User, error) {
	if name == "" {
		s.Logger.Error("User name is empty")
		return nil, ErrUserNameEmpty
	}

	user, err := s.Storage.GetUser(name)
	if err != nil {
		s.Logger.Error("Failed to get user", zap.String("name", name), zap.Error(err))
		return nil, err
	}

	return user, nil
}

// AddUserLocation adds a new location to a user's list of locations and logs the operation
// Returns an error if the user does not exist or if the user name is empty
func (s *UserService) AddUserLocation(userName, locationName, country string) error {
	if userName == "" {
		s.Logger.Error("User name is empty")
		return ErrUserNameEmpty
	}

	if err := s.Storage.AddLocation(userName, locationName, country); err != nil {
		s.Logger.Error("Failed to add location to user", zap.String("user", userName), zap.String("location", locationName), zap.Error(err))
		return err
	}

	s.Logger.Debug("Added location to user", zap.String("user", userName), zap.String("location", locationName))
	return nil
}

// AddUserTrip adds a new trip to a user's list of trips and logs the operation
// Returns an error if the user does not exist or if the user name is empty
func (s *UserService) AddUserTrip(userName, tripName string, locationNames []string) error {
	if userName == "" {
		s.Logger.Error("User name is empty")
		return ErrUserNameEmpty
	}

	if err := s.Storage.AddTrip(userName, tripName, locationNames); err != nil {
		s.Logger.Error("Failed to add trip to user", zap.String("user", userName), zap.String("trip", tripName), zap.Strings("locations", locationNames), zap.Error(err))
		return err
	}

	s.Logger.Debug("Added trip to user", zap.String("user", userName), zap.String("trip", tripName), zap.Strings("locations", locationNames))
	return nil
}
