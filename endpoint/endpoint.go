package endpoint

import (
	"log"
	"net/http"
	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/router"
)

type Endpoint struct {
	config *config.Config
	router *router.Router
}

func (endpoint Endpoint) handler(writer http.ResponseWriter, request *http.Request) {
	endpoint.router.Route(request)
}

func NewEndpoint(config *config.Config, router *router.Router) Endpoint {
	return Endpoint{config, router}
}

func (endpoint Endpoint) Start() {
	log.Printf("Starting endpoint on port%s\n", endpoint.config.Port)
	http.HandleFunc(endpoint.config.Route, endpoint.handler)
	log.Fatal(http.ListenAndServe(endpoint.config.Port, nil))
}
