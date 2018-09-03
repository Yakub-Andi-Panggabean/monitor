package repository

import (
	"firstwap.com/sms-monitoring-api/entity"
)

type CdrRepsitory interface {
	FetchLatestUpdatedCDR() entity.CDR
}
