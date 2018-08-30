package repository

import (
	"firstwap.com/sms-monitoring-api/entity"
)

type CdrRepsitory interface {
	FetchCDR(limit int) entity.CDR
}
