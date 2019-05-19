package main

import "fmt"

// simple initializer that always returns a hard-coded message
func NewMessage() Message {
	return Message("Hi there!")
}

type Message string

// an initializer for Greeter as well
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	// using the dependency injection design principle.
	// downside to dependency injection is the need for so many initialization steps.
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}
