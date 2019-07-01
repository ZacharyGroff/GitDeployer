package router

import (
	"github.com/ZacharyGroff/GitHooks/processors"
	"github.com/ZacharyGroff/GitHooks/models"
	"github.com/ZacharyGroff/GitHooks/validation"
)

type Router struct {
	config *config.Config
	pushProcessor *processors.PushProcessor
	validator *validation.Validator
}

func (router Router) Route(message *models.Message) {
		if !router.validate(message) {
			log.Fatal("Message validation failed.")
		}
		
		err := router.routeEvent(message)
		if err != nil {
			log.Fatal("Unable to route event with error %s\n", err)
		}
}

func (router Router) validate(message *models.Message) error {
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
