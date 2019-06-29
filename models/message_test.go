package models

import (
	"testing"
	"strings"
	"net/http/httptest"
)

func TestGetHeaderFieldSuccess(t *testing.T) {
	expected := "push"
	
	request := httptest.NewRequest("POST", "/", strings.NewReader("{}"))
	request.Header.Add("X-Github-Event", "push")

	message, _ := NewMessage(request)		
	actual, _ := message.GetHeaderField("X-Github-Event")

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestGetHeaderFieldErrorNonexistant(t *testing.T) {
	request := httptest.NewRequest("POST", "/", strings.NewReader("{}"))
	request.Header.Add("X-Github-Event", "push")

	message, _ := NewMessage(request)
	_ , err := message.GetHeaderField("X-Github-Event")

	if err != nil {
		t.Errorf("Did not expect error from GetHeaderField")
	}
}

func TestGetHeaderFieldErrorExists(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected error from GetHeaderField")
		}
	}()

	request := httptest.NewRequest("POST", "/", strings.NewReader("{}"))
	
	message, _ := NewMessage(request)
	message.GetHeaderField("X-Github-Event")
}
