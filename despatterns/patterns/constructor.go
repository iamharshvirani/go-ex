package patterns

import "fmt"

// Object is a struct we want to create.
type Object struct {
	Name string
}

// NewObject is the constructor function for Object.
func NewObject(name string) *Object {
	return &Object{Name: name}
}

func (o *Object) Greet() {
	fmt.Printf("Hello, I am %s\n", o.Name)
}

func RunConstructor() {
	fmt.Println("--- Constructor Pattern ---")
	obj := NewObject("Constructor")
	obj.Greet()
}
