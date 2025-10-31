package models

type LocationName string

type Location struct {
	Name    LocationName
	Country string
}

type Locations map[LocationName]Location
