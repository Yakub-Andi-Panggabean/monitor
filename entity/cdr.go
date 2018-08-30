package entity

import (
	"time"
)

type CDR struct {
	SubmmitedDateTime time.Time
	QueueMessageID    string
	DateTime          time.Time
}
