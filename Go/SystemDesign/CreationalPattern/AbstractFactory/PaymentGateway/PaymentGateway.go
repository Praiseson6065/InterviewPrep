package paymentgateway



type ServiceFactory interface{
	CreatePaymentGateway() PaymentGateway
	CreateNotificationService() NotificationService
}

type PaymentGateway interface{
	ProcessPayment(amount float64)
}

type NotificationService interface {
	Send(msg string)
}


func GetFactory(country string) ServiceFactory {
	switch country {

	case "INDIA":
		return &IndiaFactory{}

	case "US":
		return &USFactory{}

	default:
		return nil
	}
}


// func main() {

	// factory := GetFactory("INDIA")

	// payment := factory.CreatePaymentGateway()
	// notification := factory.CreateNotificationService()

	// payment.ProcessPayment(1000)

	// notification.Send("Payment Successful")

// }