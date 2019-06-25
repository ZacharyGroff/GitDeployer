package processors

import (
	"log"
	"os/exec"
	"github.com/ZacharyGroff/GitHooks/config"
)

type Deployer struct {
	config *config.Config
}

func NewDeployer(config *config.Config) *Deployer {
	return &Deployer{config} 
}

func (deployer Deployer) Deploy() string {
	out, err := exec.Command(deployer.config.ScriptPath).Output()
	
	if err != nil {
		log.Fatalf("Error while deploying:\n%s\n", err)
	}

	return string(out)	
}
