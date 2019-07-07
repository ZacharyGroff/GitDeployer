//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/GitDeployer/config"
	"github.com/ZacharyGroff/GitDeployer/processors"
	"github.com/ZacharyGroff/GitDeployer/validation"
	"github.com/ZacharyGroff/GitDeployer/endpoint"
	"github.com/ZacharyGroff/GitDeployer/router"
)

func InitializeEndpoint() endpoint.Endpoint {
	wire.Build(endpoint.NewEndpoint, router.NewRouter, 
		processors.NewPushProcessor, processors.NewDeployer,
		config.NewConfig, validation.NewValidator)

	return endpoint.Endpoint{}
}
