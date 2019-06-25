//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/processors"
	"github.com/ZacharyGroff/GitHooks/validation"
	"github.com/ZacharyGroff/GitHooks/endpoint"
)

func InitializeEndpoint() endpoint.Endpoint {
	wire.Build(endpoint.NewEndpoint, processors.NewPushProcessor, processors.NewDeployer, 
		config.NewConfig, validation.NewValidator)

	return endpoint.Endpoint{}
}
