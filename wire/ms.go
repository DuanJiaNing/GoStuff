package main

import "fmt"

type Message string

func NewMessage(msg string) Message {
	return Message(msg)
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Greeter struct {
	Message Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter, count int) Event {
	return Event{Greeter: g, count: count}
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
	count   int
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	for i := e.count; i > 0; i-- {
		fmt.Println(msg)
	}
}
