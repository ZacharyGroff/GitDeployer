package validation

import (
	"os"
	"log"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"

	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/models"
)

type Validator struct {
	config *config.Config
	secret []byte
}

func (validator Validator) ValidateHMAC(message *models.Message) bool {
	if validator.config.Validate == false {
		return true
	}
	
	hash := hmac.New(sha1.New, validator.secret)
	_, err := hash.Write([]byte(message.Body))

	if err != nil {
		log.Printf("Error creating the message hash. Error: %s", err)
		return false
	}

	hexHash := hex.EncodeToString(hash.Sum(nil))	
	log.Printf("hash: %s", hexHash)
	return hmac.Equal([]byte(hexHash), message.Hmac)
}

func NewValidator(config *config.Config) *Validator {
	return &Validator{config, []byte(os.Getenv("GITHOOKS_SECRET"))}
}
