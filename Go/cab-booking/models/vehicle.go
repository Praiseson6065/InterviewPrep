package models


type VehicleType string

const (
	VehicleBike VehicleType="BIKE"
	VehicleSedanCar VehicleType="SEDANCAR"
	VehicleMiniCar VehicleType ="MINICAR"
	VehiclePremierCar VehicleType="PRIMIERCAR"
	VehicleAuto VehicleType = "AUTO"

)

type Vehicle struct{

	VehicleNo string
	VehicleType VehicleType
	Model string


}