package validation

import (
	"testing"
	"github.com/ZacharyGroff/GitDeployer/config"
)

func TestValidateHMACPasses(t *testing.T) {
	conf := &config.Config{Validate: true}
	validator := Validator{secret: []byte("test-secret"), config: conf}
	body := []byte("test message")
	hash := []byte("e67a645c28f5decb9074052d12940ea248d34e2c")

	err := validator.ValidateHmac(hash, body)
	if err != nil {
		t.Error("Expected: true\tActual: false")
	}
}

func TestValidateHMACFails(t *testing.T) {
	conf := &config.Config{Validate: true}
	validator := Validator{secret: []byte("test-secret"), config: conf}
	body := []byte("test message")
	hash := []byte("5058f31ea908a1c5f96be7a02d08eb91fa2f467e")

	err := validator.ValidateHmac(hash, body)
	if err == nil {
		t.Error("Expected: true\tActual: false")
	}
}
