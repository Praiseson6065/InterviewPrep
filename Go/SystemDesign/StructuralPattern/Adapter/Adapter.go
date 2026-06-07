package adapter

import "fmt"

type PaymentGateway interface {
	ProcessPayment(amount float64)
}

type PayUPaymentGateway struct{

}

func (p *PayUPaymentGateway) ProcessPayment(amount float64){
	fmt.Printf("Processing payment ₹%.2f using PayUPaymentGateway\n", amount)
}

type RazorpayApi struct {
	
}

func (r *RazorpayApi) MakePayment(amount float64){
	fmt.Printf("Processing payment ₹%.2f using RazorpayAPI\n", amount)
}

type RazorpayPaymentGatewayAdapter struct{
	Sdk *RazorpayApi
}


func (p *RazorpayPaymentGatewayAdapter) ProcessPayment(amount float64){
	fmt.Printf("Processing payment ₹%.2f using RazorpayGatewayAdapter\n", amount)
	p.Sdk.MakePayment(amount)
}


// func main()  {

// 	var gateway adapter.PaymentGateway

// 	gateway = &adapter.RazorpayPaymentGatewayAdapter{
// 		Sdk: &adapter.RazorpayApi{},
// 	}

// 	gateway.ProcessPayment(1000)
	
// }




