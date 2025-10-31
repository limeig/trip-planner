package models

type User struct {
	Name      string
	Trips     []Trip
	Locations Locations
}
