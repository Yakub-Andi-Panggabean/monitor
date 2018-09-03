package usecase

import (
	"firstwap.com/sms-monitoring-api/entity"
)

//CDRUsecase represent usecase of cdr
type CDRUsecase interface {
	FindDeliveryMetric() (*entity.CDR, error)
}
