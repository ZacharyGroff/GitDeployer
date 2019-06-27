package models

import (
	"testing"
	"strings"
	"net/http/httptest"
)

func TestGetHeaderKeySuccess(t *testing.T) {
	expected := "testEvent"
	
	request := httptest.NewRequest("POST", "/", strings.NewReader("testBody"))
	request.Header.Add("X-Github-Event", "testEvent")

	message := NewMessage(request)		
	actual, err := message.GetHeaderKey("X-Github-Event")

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestGetHeaderKeyErrorNonexistant(t *testing.T) {
	request := httptest.NewRequest("POST", "/", strings.NewReader("testBody"))
	request.Header.Add("X-Github-Event", "testEvent")

	message := NewMessage(request)
	actual, err := message.GetHeaderKey("X-Github-Event")

	if err != nil {
		t.Errorf("Did not expect error from GetHeaderKey")
	}
}

func TestGetHeaderKeyErrorExists(t *testing.T) {
	request := httptest.NewRequest("POST", "/", strings.NewReader("testBody"))
	
	message := NewMessage(request)
	actual, err := message.GetHeaderKey("X-Github-Event")

	if err == nil {
		t.Errorf("Expected error from GetHeaderKey")
	}
}
