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

func (message Message) GetHeaderField(key string) (string, Error) {
	field := message.Header.get(key)

	if strings.Compare(field, "") == 0 {
		err := fmt.Errorf("Header Field %s not found", key)
		return nil, err
	}

	return field, nil
}

func (message Message) setPayload(body []byte) Error {
	event, err := message.GetHeaderField("X-Github-Event")

	if err != nil {
		return err
	}

	switch event {
	case "push":
		message.Payload = NewPushPayload(body)
	default:
		err := fmt.Errorf("Unhandled event %s detected.", event)
		return err
	}
}

func NewMessage(request *http.Request) (*Message, Error) { 
	message := Message{Header: request.Header}
	body := ioutil.ReadAll(request.Body)
	err := message.setPayload(body)
	
	if err != nil {
		return nil, err
	}

	return &message, nil
}
