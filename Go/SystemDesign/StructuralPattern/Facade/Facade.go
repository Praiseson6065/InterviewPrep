package main

import "fmt"

type PayementService struct{}

func (p *PayementService) Pay(amount float64) bool{
	fmt.Printf("Payment Processing %.2f\n",amount)
	return true

}

type NotificationService struct{}

func (n *NotificationService) Send(message string){
	fmt.Printf("Notification Sent: %s\n",message)
}

type InventoryService struct{}

func (i *InventoryService) CheckStock(product string) bool {
	fmt.Println("Checking inventory for:", product)
	return true
}


type OrderFacade struct{
	inventory *InventoryService
	payment *PayementService
	notification *NotificationService
}

func NewOrderFacade() *OrderFacade{
	return &OrderFacade{
		inventory: &InventoryService{},
		payment: &PayementService{},
		notification: &NotificationService{},
	}
}


func (o *OrderFacade) PlaceOrder(product string,amount float64){
	if !o.inventory.CheckStock(product) {
		fmt.Println("Out of stock")
		return
	}

	if !o.payment.Pay(amount) {
		fmt.Println("Payment failed")
		return
	}


	o.notification.Send("Order placed successfully",)
}

func main(){

	orderFacade :=NewOrderFacade()

	orderFacade.PlaceOrder("iPhone 17",99999,)



}
