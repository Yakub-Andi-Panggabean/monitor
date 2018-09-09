package domain

import (
	"time"
)

//CDR struct which took take three dates variable from CDR table
type CDR struct {
	Provider          string
	SubmmitedDateTime time.Time
	QueueMessageID    string
	DateTime          time.Time
	DeliveryCode      string
	Interval          int64
}
