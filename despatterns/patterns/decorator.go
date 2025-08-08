package patterns

import "fmt"

// Component is the interface for objects that can have responsibilities added to them.
type Component interface {
	Operation() string
}

// ConcreteComponent is an object to which additional responsibilities can be attached.
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "I am a concrete component"
}

// Decorator maintains a reference to a Component object and defines an interface
// that conforms to Component's interface.
type Decorator struct {
	component Component
}

func (d *Decorator) Operation() string {
	return d.component.Operation()
}

// ConcreteDecoratorA adds responsibilities to the component.
type ConcreteDecoratorA struct {
	Decorator
}

func NewConcreteDecoratorA(component Component) *ConcreteDecoratorA {
	return &ConcreteDecoratorA{Decorator{component: component}}
}

func (d *ConcreteDecoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", d.Decorator.Operation())
}

// ConcreteDecoratorB adds other responsibilities.
type ConcreteDecoratorB struct {
	Decorator
}

func NewConcreteDecoratorB(component Component) *ConcreteDecoratorB {
	return &ConcreteDecoratorB{Decorator{component: component}}
}

func (d *ConcreteDecoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)", d.Decorator.Operation())
}

func RunDecorator() {
	fmt.Println("--- Decorator Pattern ---")
	component := &ConcreteComponent{}
	decoratorA := NewConcreteDecoratorA(component)
	decoratorB := NewConcreteDecoratorB(decoratorA)

	fmt.Println(decoratorB.Operation())
}
