package entity

import (
	"time"
)

//CDR struct which took take three dates variable from CDR table
type CDR struct {
	SubmmitedDateTime time.Time
	QueueMessageID    string
	DateTime          time.Time
	Interval          int64
}
