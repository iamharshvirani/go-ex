package patterns

import "fmt"

// IProduct defines the interface for products.
type IProduct interface {
	GetName() string
}

// ConcreteProductA is one implementation of IProduct.
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "Product A"
}

// ConcreteProductB is another implementation of IProduct.
type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "Product B"
}

// GetProduct is the factory method.
func GetProduct(productType string) (IProduct, error) {
	if productType == "A" {
		return &ConcreteProductA{}, nil
	}
	if productType == "B" {
		return &ConcreteProductB{}, nil
	}
	return nil, fmt.Errorf("invalid product type")
}

func RunFactory() {
	fmt.Println("--- Factory Pattern ---")
	productA, _ := GetProduct("A")
	fmt.Println("Created:", productA.GetName())

	productB, _ := GetProduct("B")
	fmt.Println("Created:", productB.GetName())
}
