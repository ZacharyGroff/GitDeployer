package validation

import (
	"os"
	"crypto/hmac"
	"crypto/sha1"

	"github.com/ZacharyGroff/GitHooks/config"
)

type Validator struct {
	config *config.Config
	secret []byte
}

func (v Validator) ValidateHMAC(message []byte, messageHMAC []byte) bool {
	if v.config.Validate == false {
		return true
	}
	
	hash := hmac.New(sha1.New, v.secret)
	hash.Write([]byte(message))
	
	return hmac.Equal(hash.Sum(nil), messageHMAC)
}

func NewValidator(config *config.Config) *Validator {
	return &Validator{config, []byte(os.Getenv("GITHOOKS_SECRET"))}
}
