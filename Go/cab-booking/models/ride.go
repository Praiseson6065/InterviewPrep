package models

import "time"

type RideStatus string

const(
	Requested RideStatus="REQUESTED" 
	Accepted RideStatus="ACCEPTED" 
	InProgress RideStatus="IN_PROGRESS" 
	Completed RideStatus="COMPLETED" 
	Cancelled RideStatus="CANCELLED" 

)

type ride struct{
	Id string
	RiderId string
	DriverId *string
	PickupLocation Location
	DropLocation Location
	VehicleType VehicleType
	Status RideStatus
	OTP string
	Fare *Fare
	RequestedAt time.Time
	AcceptedAt time.Time
	StartedAt time.Time
	CompletedAt time.Time
	Payment 
}