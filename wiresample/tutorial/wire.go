//+build wireinject

// https://godoc.org/go/build#hdr-Build_Constraints
package main

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
