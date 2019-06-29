package models

import (
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Message struct {
	Header http.Header
	Payload Payload
}

func (message Message) GetHeaderField(key string) (string, error) {
	field := message.Header.Get(key)

	if strings.Compare(field, "") == 0 {
		err := fmt.Errorf("Header Field %s not found", key)
		return field, err
	}

	return field, nil
}

func (message Message) setPayload(body []byte) error {
	event, err := message.GetHeaderField("X-Github-Event")

	if err != nil {
		return err
	}

	switch event {
	case "push":
		message.Payload, err = NewPushPayload(body)
	default:
		err = fmt.Errorf("Unhandled event %s detected.", event)
	}

	if err != nil {
		return err
	}

	return nil
}

func NewMessage(request *http.Request) (*Message, error) { 
	message := Message{Header: request.Header}
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return nil, err
	}

	err = message.setPayload(body)
	
	if err != nil {
		return nil, err
	}

	return &message, nil
}
