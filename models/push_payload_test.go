package models

import (
	"os"
	"testing"
	"io/ioutil"
)

func TestNewPayloadSuccess(t *testing.T) {
	file, _ := os.Open("push_payload_test.json")
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	
	_ , err := NewPushPayload(byteValue)

	if err != nil {
		t.Errorf("Unable to create PushPayload with error: %s", err.Error())
	}
}
