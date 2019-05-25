package main

import (
	"errors"
	"fmt"
	"time"
)

// simple initializer that always returns a hard-coded message
func NewMessage(phrase string) Message {
	return Message(phrase)
}

type Message string

// an initializer for Greeter as well
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

type Greeter struct {
	Message Message
	Grumpy  bool
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// add to return error
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
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
	//message := NewMessage()
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	//
	//event.Start()

	e, err := InitializeEvent("hello")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
	}
	e.Start()
}
