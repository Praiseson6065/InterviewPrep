package paymentgateway

import "fmt"


type IndiaFactory struct{

}

func (i *IndiaFactory) CreatePaymentGateway()PaymentGateway{

	return Razorpay{}

}

func (i *IndiaFactory) CreateNotificationService()NotificationService{
	return SMSService{}
}

type Razorpay struct {
}

func (r Razorpay) ProcessPayment(amt float64){
	fmt.Printf("Processing ₹%.2f using Razorpay\n", amt)
}

type SMSService struct{

}

func (s SMSService) Send(msg string){
	fmt.Println("SMS:", msg)
}
