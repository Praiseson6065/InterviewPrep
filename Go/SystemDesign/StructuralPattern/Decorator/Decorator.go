package main

import "fmt"


type Coffee interface {
	GetDescription() string
	GetCost() float64
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) GetDescription() string{
	return "Coffee : Simple Coffee"
}

func (c *SimpleCoffee) GetCost() float64{
	return 36.5
}

type CoffeeDecorator struct{
	coffee Coffee
}

type Milk struct{
	CoffeeDecorator
}

func NewMilk(coffee Coffee) Coffee {
	return &Milk{
		CoffeeDecorator: CoffeeDecorator{
			coffee: coffee,
		},
	}
}


func (m *Milk) GetDescription() string{
	return m.coffee.GetDescription() + ", Milk"
}

func (m *Milk) GetCost() float64{
	return m.coffee.GetCost()+20.5
}


type Sugar struct{
	CoffeeDecorator
}

func NewSugar(Coffee Coffee)Coffee{
	return &Sugar{
		CoffeeDecorator: CoffeeDecorator{
			coffee : Coffee, 
		},
	}
}

func (s *Sugar) GetDescription() string{

	return s.coffee.GetDescription() + ", Sugar"
}


func (s *Sugar) GetCost() float64{

	return s.coffee.GetCost()+10.33
}


type Caramel struct {
	CoffeeDecorator
}

func NewCaramel(coffee Coffee) Coffee {
	return &Caramel{
		CoffeeDecorator: CoffeeDecorator{
			coffee: coffee,
		},
	}
}

func (c *Caramel) GetDescription() string {
	return c.coffee.GetDescription() + ", Caramel"
}

func (c *Caramel) GetCost() float64 {
	return c.coffee.GetCost() + 30
}


func main() {

	var coffee Coffee = &SimpleCoffee{}

	coffee = NewMilk(coffee)
	coffee = NewSugar(coffee)
	coffee = NewCaramel(coffee)

	fmt.Println("Description:", coffee.GetDescription())
	fmt.Printf("Total Cost: ₹%.2f\n", coffee.GetCost())
}