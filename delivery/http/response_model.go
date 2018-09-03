package http

import (
	"time"
)

type Metric struct {
	Provider          string    `json:"provider"`
	SmsAPIMessageID   string    `json:"messageId"`
	AcceptedDateTime  time.Time `json:"acceptedDate"`
	DeliveredDateTime time.Time `json:"deliveredDate"`
	Interval          int64     `json:"interval"`
}

//MetricsRespose list of metrics which will be displayed by api
type MetricsRespose struct {
	Metrics []Metric `json:"metrics"`
}
