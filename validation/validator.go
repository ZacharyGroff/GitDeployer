package validation

import (
	"os"
	"fmt"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/ZacharyGroff/GitHooks/config"
)

type Validator struct {
	config *config.Config
	secret []byte
}

func (v Validator) ValidateHmac(messageHmac []byte, body []byte) error {
	if v.config.Validate == false {
		return nil
	}
	
	hash := hmac.New(sha1.New, v.secret)
	_, err := hash.Write([]byte(body))

	if err != nil {
		return err	
	}

	hexHash := hex.EncodeToString(hash.Sum(nil))	
	if hmac.Equal([]byte(hexHash), messageHmac) {
		return nil
	} else {
		return fmt.Errorf("Hmac failed validation")
	}
}

func NewValidator(config *config.Config) *Validator {
	return &Validator{config, []byte(os.Getenv("GITHOOKS_SECRET"))}
}
