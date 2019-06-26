package models

import (
	"testing"
	"strings"
	"net/http/httptest"
)

func TestParseRequestEvent(t *testing.T) {
	expected := "testEvent"
	
	request := httptest.NewRequest("POST", "/", strings.NewReader("testBody"))
	request.Header.Add("X-Github-Event", "testEvent")
	request.Header.Add("X-Hub-Signature", "sha1=test")
	
	message := Message{}
	parseRequest(&message, request)
	actual := message.Event

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestParseRequestBody(t *testing.T) {
	expected := "testBody"
	
	request := httptest.NewRequest("POST", "/", strings.NewReader("testBody"))
	request.Header.Add("X-Github-Event", "testEvent")
	request.Header.Add("X-Hub-Signature", "sha1=test")
	
	message := Message{}
	parseRequest(&message, request)
	actual := string(message.Body)	

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestParseRequestHmac(t *testing.T) {
	expected := "testHmac"
	
	request := httptest.NewRequest("POST", "/", strings.NewReader("testBody"))
	request.Header.Add("X-Github-Event", "testEvent")
	request.Header.Add("X-Hub-Signature", "sha1=testHmac")
	
	message := Message{}
	parseRequest(&message, request)
	actual := string(message.Hmac)	

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
