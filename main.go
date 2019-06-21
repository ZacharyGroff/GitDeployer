package main

func main() {
	configPtr := NewConfig() 
	endpoint := Endpoint{configPtr}
	endpoint.Start()
}
