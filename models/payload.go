package models

type Payload interface {
	parsePayload([]byte) error
}
