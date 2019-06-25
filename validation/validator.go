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

func (v Validator) ValidateHMAC(message []byte, messageHMAC []byte) bool {
	if v.config.Validate == false {
		return true
	}
	
	hash := hmac.New(sha1.New, v.secret)
	_, err := hash.Write([]byte(message))

	if err != nil {
		log.Printf("Error creating the message hash. Error: %s", err)
		return false
	}

	hexHash := hex.EncodeToString(hash.Sum(nil))	
	log.Printf("hash: %s", hexHash)
	return hmac.Equal([]byte(hexHash), messageHMAC)
}

func NewValidator(config *config.Config) *Validator {
	return &Validator{config, []byte(os.Getenv("GITHOOKS_SECRET"))}
}
