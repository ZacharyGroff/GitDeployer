package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type Endpoint struct {
	config *Config
}

func handler(writer http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	event := request.Header["X-Github-Event"][0]
	
	fmt.Printf("Headers:\n%v\n\n", request.Header)
	fmt.Printf("Body:\n%s\n\n", body)
	fmt.Printf("X-GitHub-Event: %s\n", event)

	routeEvent(event, request)
}

func routeEvent(event string, request *http.Request) {
	switch event {
	case "push":
		fmt.Println("Push event detected.")
	default:
		log.Printf("Unhandled event %s detected.\n", event)
	}
}

func (endpoint Endpoint) Start() {
	log.Printf("Starting endpoint on port%s\n", endpoint.config.Port)
	http.HandleFunc(endpoint.config.Route, handler)
	log.Fatal(http.ListenAndServe(endpoint.config.Port, nil))
}
