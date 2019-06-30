package processors

import (
	"github.com/ZacharyGroff/GitHooks/models"
)

type Processor interface {
	HandleMessage(*models.Message)
}
