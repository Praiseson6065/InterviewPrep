package paymentgateway

import "fmt"


type USFactory struct{

}

func (u *USFactory) CreatePaymentGateway()PaymentGateway{

	return Paypal{}

}

func (u *USFactory) CreateNotificationService()NotificationService{
	return EmailService{}
}

type Paypal struct {
}

func (p Paypal) ProcessPayment(amt float64){
	fmt.Printf("Processing ₹%.2f using Paypal\n", amt)
}

type EmailService struct{

}

func (e EmailService) Send(msg string){
	fmt.Println("Email:", msg)
}
