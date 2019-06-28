package endpoint

import (
	"log"
	"net/http"
	"io/ioutil"

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
	message, err := models.NewMessage(request)

	if err != nil {
		log.Fatalf("Failed to create message with error: %s\n", err)
	}

	hmac, _ := message.GetHeaderField("X-Hub-Signature")
	trimmedHmac := []byte(hmac)[5:]
	body, _ := ioutil.ReadAll(request.Body)

	if endpoint.validator.ValidateHMAC(trimmedHmac, body) {
		err := endpoint.routeEvent(message)

		if err != nil {
			log.Fatalf("Unable to route message")
		}
	} else {
		log.Fatalf("Message validation failed. Message HMAC: %s\n", hmac)
	}
}

func (endpoint Endpoint) routeEvent(message *models.Message) error {
	event, err := message.GetHeaderField("X-Github-Event")

	if err != nil {
		return err
	}

	switch event {
	case "push":
		endpoint.pushProcessor.HandleMessage(message)
	default:
		log.Printf("Unhandled event %s detected.\n", event)
	}

	return nil
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
