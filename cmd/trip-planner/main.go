package main

import (
	"fmt"

	"trip-planner/internal/storage"
)

func main() {
	fmt.Println("Hello, World")

	// Create a repository
	repository := storage.New()

	repository.AddUser("Alina")

	repository.AddLocation("Alina", "Izumo", "Japan")
	repository.AddLocation("Alina", "Tokyo", "Japan")

	repository.AddTrip("Alina", "Japan", []string{"Izumo", "Osaka", "Tokyo"})
	user, _ := repository.GetUser("Alina")

	fmt.Printf("Get user %s with trips %v\n", user.Name, user.Trips)
}
