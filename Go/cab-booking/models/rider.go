package models

import "time"


type Rider struct{
	Id string
	Name string
	Email string
	RideHistory []string
	CreatedAt time.Time
}