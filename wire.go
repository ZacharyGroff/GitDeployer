//+build wireinject

package main

import (
	"github.com/google/wire"
	"config"
	"processors"
)

func InitializeEndpoint() Endpoint {
	wire.Build(NewEndpoint, processors.NewPushProcessor, processors.NewDeployer, config.NewConfig)
	return Endpoint{}
}
