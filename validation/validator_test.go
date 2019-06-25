package validation

import (
	"testing"

	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/models"
)

func TestValidateHMACPasses(t *testing.T) {
	conf := &config.Config{Validate: true}
	validator := Validator{secret: []byte("test-secret"), config: conf}
	body := []byte("test message")
	hash := []byte("e67a645c28f5decb9074052d12940ea248d34e2c")

	message := models.Message{Body: body, Hmac: hash}
	if validator.ValidateHMAC(&message) == false {
		t.Error("Expected: true\nActual: false\n")
	}
}

func TestValidateHMACFails(t *testing.T) {
	conf := &config.Config{Validate: true}
	validator := Validator{secret: []byte("test-secret"), config: conf}
	body := []byte("test message")
	hash := []byte("5058f31ea908a1c5f96be7a02d08eb91fa2f467e")

	message := models.Message{Body: body, Hmac:hash}
	if validator.ValidateHMAC(&message) == true {
		t.Error("Expected: false\nActual: true\n")
	}
}
