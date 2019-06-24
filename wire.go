//+build wireinject

package main

import (
	"github.com/google/wire"
	"config"
)

func InitializeEndpoint() Endpoint {
	wire.Build(NewEndpoint, config.NewConfig)
	return Endpoint{}
}
