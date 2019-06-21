import (
	"encoding/json"
	"os"
	"fmt"
	"log"
)

type Configuration struct {
	port string
	route string
	validate bool
}

func parseConfig() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)	
	conf := Configuration{}
	err := decoder.Decode(&conf)
	
	return conf, err
}

func NewConfiguration() Configuration {
	conf, err = parseConfig()
	
	if err != nil {
		log.Printf("Error while parsing config")
	}
	
	return conf
}
