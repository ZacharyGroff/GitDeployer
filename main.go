package main

func main() {
	//accept as user input or read from file in future
	port := uint16(8080)
	route := "/webhook"

	endpoint := Endpoint{port, route}
	endpoint.Start()
}
