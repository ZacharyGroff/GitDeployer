package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/processors"
)

type Endpoint struct {
	config *config.Config
	pushProcessor *processors.PushProcessor
}

func (endpoint Endpoint) handler(writer http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	event := request.Header["X-Github-Event"][0]
	
	fmt.Printf("Headers:\n%v\n\n", request.Header)
	fmt.Printf("Body:\n%s\n\n", body)
	fmt.Printf("X-GitHub-Event: %s\n", event)

	endpoint.routeEvent(event, request)
}

func (endpoint Endpoint) routeEvent(event string, request *http.Request) {
	switch event {
	case "push":
		endpoint.pushProcessor.HandleRequest(request)
	default:
		log.Printf("Unhandled event %s detected.\n", event)
	}
}

func NewEndpoint(config *config.Config, pushProcessor *processors.PushProcessor) Endpoint {
	return Endpoint{config, pushProcessor}
}

func (endpoint Endpoint) Start() {
	log.Printf("Starting endpoint on port%s\n", endpoint.config.Port)
	http.HandleFunc(endpoint.config.Route, endpoint.handler)
	log.Fatal(http.ListenAndServe(endpoint.config.Port, nil))
}
