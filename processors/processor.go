package processors

import (
	"github.com/ZacharyGroff/GitDeployer/models"
)

type Processor interface {
	HandleMessage(*models.Message)
}
