package router

import (
	"log"
	"net/http"
	"io/ioutil"
	"github.com/ZacharyGroff/GitHooks/processors"
	"github.com/ZacharyGroff/GitHooks/models"
	"github.com/ZacharyGroff/GitHooks/validation"
)

type Router struct {
	pushProcessor *processors.PushProcessor
	validator *validation.Validator
}

func (router Router) Route(request *http.Request) {
	message, err := models.NewMessage(request)

	if err != nil {
		log.Fatalf("Failed to create message with error: %s\n", err)
	}

	if !router.validate(message, request) {
		log.Fatal("Message validation failed.")
	}

	err = router.routeEvent(message)
	if err != nil {
		log.Fatalf("Unable to route event with error %s\n", err)
	}
}

func (router Router) validate(message *models.Message, request *http.Request) bool {
	hmac, _ := message.GetHeaderField("X-Hub-Signature")
	trimmedHmac := []byte(hmac)[5:]
	body, _ := ioutil.ReadAll(request.Body)

	return router.validator.ValidateHMAC(trimmedHmac, body)
}

func (router Router) routeEvent(message *models.Message) error {
	event, err := message.GetHeaderField("X-Github-Event")

	if err != nil {
		return err
	}

	switch event {
	case "push":
		router.pushProcessor.HandleMessage(message)
	default:
		log.Printf("Unhandled event %s detected.\n", event)
	}

	return nil
}

func NewRouter(p *processors.PushProcessor, v *validation.Validator) *Router {
	return &Router{p, v}
}
