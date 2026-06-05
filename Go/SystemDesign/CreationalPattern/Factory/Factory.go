package main

import "fmt"

type PaymentGateway interface{
	ProcessPayment(amount float64) error
}

type PaymentGatewayType string 

const (
	UPIGateway PaymentGatewayType = "UPI"
	CardGateway   PaymentGatewayType = "CARD"
)


type UpiPayment struct{}

func (u *UpiPayment)ProcessPayment(amount float64) error {

	fmt.Printf("Processing ₹%.2f via UPI\n",amount)
	return nil
}

type CardPayment struct{}
func(c *CardPayment)ProcessPayment(amount float64)error{
	fmt.Printf("Processing ₹%.2f via Card\n",amount)
	return nil
}

func GetPaymentGateway(gateway PaymentGatewayType)(PaymentGateway,error){
	switch gateway {
	case UPIGateway:
		return &UpiPayment{},nil
	case CardGateway:
		return &CardPayment{},nil
	default : 
		return nil,fmt.Errorf("unsupported gateway: %s", gateway)
	}

}

type PaymentService struct{}

func (ps *PaymentService) MakePayment(
	gatewayType PaymentGatewayType,
	amount float64,
) error {

	gateway, err := GetPaymentGateway(gatewayType)
	if err != nil {
		return err
	}

	return gateway.ProcessPayment(amount)
}

// func main() {

// 	service := &PaymentService{}

// 	service.MakePayment(UPIGateway, 1000)
// 	service.MakePayment(CardGateway, 23000)
	
// }

