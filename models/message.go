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

func (message Message) FromRequest(request *http.Request) {
	fullHmac := request.Header["X-Hub-Signature"][0]
	
	message.Event = request.Header["X-Github-Event"][0]
	message.Body, _ = ioutil.ReadAll(request.Body)
	message.Hmac = []byte(fullHmac[5:])
}
