package main

import (
	"fmt"
	"log"
	"net/http"
)

func webhookHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Headers:\n%v\n\n", request.Header)
	fmt.Printf("Body:\n%v\n", request.Body)
}

func main() {
	//accept as user input or read from file in future
	port := 8080
	portString := fmt.Sprintf(":%d", port)
	log.Printf("Starting endpoint on port %d\n", port)
	
	http.HandleFunc("/webhook", webhookHandler)
	log.Fatal(http.ListenAndServe(portString, nil))
}
