package endpoint

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/processors"
	"github.com/ZacharyGroff/GitHooks/validation"
	"github.com/ZacharyGroff/GitHooks/models"
)

type Endpoint struct {
	config *config.Config
	pushProcessor *processors.PushProcessor
	validator *validation.Validator
}


func (endpoint Endpoint) handler(writer http.ResponseWriter, request *http.Request) {
	message := models.NewMessage(request)
	
	fmt.Printf("Body:\n%s\n\n", message.Body)
	fmt.Printf("X-GitHub-Event: %s\n", message.Event)
	fmt.Printf("X-Hub-Signature: %s\n", message.Hmac)

	if endpoint.validator.ValidateHMAC(message) {
		endpoint.routeEvent(message)
	} else {
		log.Fatalf("Message validation failed. Message HMAC: %s\n", message.Hmac)
	}
}

func (endpoint Endpoint) routeEvent(message *models.Message) {
	switch message.Event {
	case "push":
		endpoint.pushProcessor.HandleMessage(message)
	default:
		log.Printf("Unhandled event %s detected.\n", message.Event)
	}
}

func NewEndpoint(
	config *config.Config, 
	pushProcessor *processors.PushProcessor, 
	validator *validation.Validator) Endpoint {
	return Endpoint{config, pushProcessor, validator}
}

func (endpoint Endpoint) Start() {
	log.Printf("Starting endpoint on port%s\n", endpoint.config.Port)
	http.HandleFunc(endpoint.config.Route, endpoint.handler)
	log.Fatal(http.ListenAndServe(endpoint.config.Port, nil))
}