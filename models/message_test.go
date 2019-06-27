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

	message := NewMessage(request.Header, request.Body)		
	actual := message.GetHeaderKey("X-Github-Event")

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestGetHeaderKeyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetHeaderKey did not panic.")
		}
	}		
	
	request := httptest.NewRequest("POST", "/", strings.NewReader("testBody"))
	
	message := NewMessage(request.Header, request.Body)
	actual := message.GetHeaderKey("X-Github-Event")
}
