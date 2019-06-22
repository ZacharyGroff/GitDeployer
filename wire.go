//+build wireinject

package main

import "github.com/google/wire"

func InitializeEndpoint() Endpoint {
	wire.Build(NewEndpoint, NewConfig)
	return Endpoint{}
}
