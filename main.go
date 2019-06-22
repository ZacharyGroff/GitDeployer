package main

func main() {
	config := NewConfig() 
	endpoint := Endpoint{config}
	endpoint.Start()
}
