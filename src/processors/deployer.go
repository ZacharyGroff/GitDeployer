package processors

import (
	"log"
	"os/exec"
	"config"
)

type Deployer struct {
	config *config.Config
}

func NewDeployer(config *config.Config) *Deployer {
	return &Deployer{config} 
}

func (deployer Deployer) Deploy() {
	out, err := exec.Command(deployer.config.ScriptPath).Output()
	
	if err != nil {
		log.Fatal("Error while deploying")
	}
	
	log.Printf("Successfully deployed script with output:\n%s\n", out)
}
