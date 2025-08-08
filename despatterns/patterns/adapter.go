package patterns

import "fmt"

// Target is the interface the client expects to work with.
type Target interface {
	Request() string
}

// Adaptee is the existing struct with an incompatible interface.
type Adaptee struct{}

// SpecificRequest is the method on the Adaptee that needs to be adapted.
func (a *Adaptee) SpecificRequest() string {
	return "I am the Adaptee with a specific request!"
}

// Adapter is the struct that adapts the Adaptee to the Target interface.
type Adapter struct {
	adaptee *Adaptee
}

// NewAdapter creates a new Adapter.
func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

// Request implements the Target interface by calling the Adaptee's method.
func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func RunAdapter() {
	fmt.Println("--- Adapter Pattern ---")
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)

	// The client code uses the adapter which conforms to the Target interface.
	fmt.Println(adapter.Request())
}
