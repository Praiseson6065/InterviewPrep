package models

import "time"

type Location struct{
	Lat float64
	Lon float64
	Address string
}

type Driver struct{
	Id string
	Name string
	Email string
	Vehicle Vehicle
	Rating float64
	TotalRides int
	CurrentLocation Location
	CreatedAt time.Time
}	