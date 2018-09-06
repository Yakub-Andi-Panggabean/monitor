package repository

import (
	"firstwap.com/sms-monitoring-api/entity"
)

type CdrRepository interface {
	FetchLatestUpdatedCDR() []*entity.CDR
}
