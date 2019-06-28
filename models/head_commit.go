package models

import (
	"time"
)

type HeadCommit struct {
    Added     []string  `json:"added"`
    Author    Author    `json:"author"`
    Committer Committer `json:"committer"`
    Distinct  bool      `json:"distinct"`
    Id        string    `json:"id"`
    Message   string    `json:"message"`
    Modified  []string  `json:"modified"`
    Removed   []string  `json:"removed"`
    Timestamp time.Time `json:"timestamp"`
    Url       string    `json:"url"`
}
