package models

type User struct {
	Name      string
	Trips     []Trip
	Locations map[string]Location
}
