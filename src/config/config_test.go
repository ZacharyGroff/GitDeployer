package config

import (
	"testing"
	"strings"
)

func TestParseConfigPort(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := ":42"
	actual := config.Port
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestParseConfigRoute(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := "/test"
	actual := config.Route
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestParseConfigValidate(t *testing.T) {
	config := Config{}
	config.parseConfig("conf_test.json")
	
	expected := true
	actual := config.Validate
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}
