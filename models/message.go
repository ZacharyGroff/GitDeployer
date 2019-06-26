package models

import (
	"net/http"
	"io/ioutil"
)

type Message struct {
	Event string
	Body []byte
	Hmac []byte
}

func parseRequest(message *Message, request *http.Request) *Message {
	fullHmac := request.Header["X-Hub-Signature"][0]
	
	message.Event = request.Header["X-Github-Event"][0]
	message.Body, _ = ioutil.ReadAll(request.Body)
	message.Hmac = []byte(fullHmac[5:])

	return message
}

func NewMessage(request *http.Request) *Message { 
	message := Message{}
	return parseRequest(&message, request)
}
