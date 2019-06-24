//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/processors"
)

func InitializeEndpoint() Endpoint {
	wire.Build(NewEndpoint, processors.NewPushProcessor, processors.NewDeployer, config.NewConfig)
	return Endpoint{}
}
