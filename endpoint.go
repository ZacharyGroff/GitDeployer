package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type Endpoint struct {
	port uint16
	route string
}

func handler(writer http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Printf("Headers:\n%v\n\n", request.Header)
	fmt.Printf("Body:\n%s\n", body)
}

func (endpoint Endpoint) Start() {
	portString := fmt.Sprintf(":%d", endpoint.port)
	log.Printf("Starting endpoint on port%s\n", portString)
	http.HandleFunc(endpoint.route, handler)
	log.Fatal(http.ListenAndServe(portString, nil))
}
