//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(str string, ct int) Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
