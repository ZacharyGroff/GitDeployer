package processors

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Deployer struct {
	scriptPath string
}

func NewDeployer(scriptPath string) *Deployer {
	return &Deployer{scriptPath} 
}

func (deployer Deployer) Deploy() Error {
	out, err := os.exec.Command(deyployer.scriptPath).Output()
	
	if err != nil {
		log.Fatal("Error while deploying")
	}
	
	log.Printf("Successfully deployed script with output:\n%s\n", out)
}
