package patterns

import "fmt"

// Observer is the interface for objects that should be notified of changes.
type Observer interface {
	Update(string)
}

// Subject is the interface for the object that holds the state and notifies observers.
type Subject interface {
	Register(Observer)
	Deregister(Observer)
	Notify()
}

// ConcreteSubject holds the state and a list of observers.
type ConcreteSubject struct {
	observers []Observer
	state     string
}

func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Deregister(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	s.Notify()
}

// ConcreteObserver implements the Observer interface.
type ConcreteObserver struct {
	name string
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

func (o *ConcreteObserver) Update(state string) {
	fmt.Printf("Observer %s received update with state: %s\n", o.name, state)
}

func RunObserver() {
	fmt.Println("--- Observer Pattern ---")
	subject := &ConcreteSubject{}

	observer1 := NewConcreteObserver("A")
	observer2 := NewConcreteObserver("B")

	subject.Register(observer1)
	subject.Register(observer2)

	subject.SetState("state 1")

	subject.Deregister(observer1)

	subject.SetState("state 2")
}
