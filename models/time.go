package models

import (
	"time"
	"bytes"
	"strconv"
)

type Time struct {
	time.Time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return t.Time.MarshalJSON()
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		t.Time = time.Time{}
		return nil
	}

	var seconds int64
	err := t.Time.UnmarshalJSON(data)
	if err == nil {
		return nil
	}

	trimmedBytes := string(bytes.Trim(data, `"`))
	seconds, err = strconv.ParseInt(trimmedBytes, 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(seconds, 0)
	return nil
}
