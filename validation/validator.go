package validation

import (
	"os"
	"log"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"

	"github.com/ZacharyGroff/GitHooks/config"
)

type Validator struct {
	config *config.Config
	secret []byte
}

func (v Validator) ValidateHMAC(messageHmac []byte, body []byte) bool {
	if v.config.Validate == false {
		return true
	}
	
	hash := hmac.New(sha1.New, v.secret)
	_, err := hash.Write([]byte(body))

	if err != nil {
		log.Printf("Error creating the message hash. Error: %s", err)
		return false
	}

	hexHash := hex.EncodeToString(hash.Sum(nil))	
	return hmac.Equal([]byte(hexHash), messageHmac)
}

func NewValidator(config *config.Config) *Validator {
	return &Validator{config, []byte(os.Getenv("GITHOOKS_SECRET"))}
}
