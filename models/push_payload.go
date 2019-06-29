package models

import (
	"encoding/json"
)

type PushPayload struct {
    After      string     `json:"after"`
    BaseRef    string     `json:"base_ref"`
    Before     string     `json:"before"`
    Commits    []Commits  `json:"commits"`
    Compare    string     `json:"compare"`
    Created    bool       `json:"created"`
    Deleted    bool       `json:"deleted"`
    Forced     bool       `json:"forced"`
    HeadCommit HeadCommit `json:"head_commit"`
    Pusher     Pusher     `json:"pusher"`
    Ref        string     `json:"ref"`
    Repository Repository `json:"repository"`
    Sender     Sender     `json:"sender"`
}

func (pushPayload PushPayload) parsePayload(body []byte) error {
	err := json.Unmarshal(body, &pushPayload)
	
	if err != nil {
		return err
	}
	
	return nil
}

func NewPushPayload(body []byte) (*PushPayload, error) {
	pushPayload := PushPayload{}
	err := pushPayload.parsePayload(body)
	
	if err != nil {
		return nil, err
	}

	return &pushPayload, nil
}
