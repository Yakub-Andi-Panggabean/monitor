package service

import (
	"firstwap.com/sms-monitoring-api/application"
	"firstwap.com/sms-monitoring-api/application/domain"
	"firstwap.com/sms-monitoring-api/application/logger"
	"firstwap.com/sms-monitoring-api/application/parser"
)

var log = logger.New("cdr_usecase")

type cdrService struct {
	cdrRepo application.CdrStorageFetcher
}

//New create cdr usecase instance
func New(cdrStorageFetcher application.CdrStorageFetcher) application.DeliveryMetricFinder {

	return &cdrService{
		cdrRepo: cdrStorageFetcher,
	}

}

func (r *cdrService) FindDeliveryMetric() ([]*domain.CDR, error) {

	result := make([]*domain.CDR, 0)
	cdrs := r.cdrRepo.FetchLatestUpdatedCDR()

	for _, cdr := range cdrs {

		messageIDDateTime, _ := parser.ParseQueueMessageID(cdr.QueueMessageID)
		time, err := parser.ConvertDateStringToTime(messageIDDateTime)

		if err != nil {
			return nil, err
		}

		cdr.Interval = int64(time.Sub(cdr.DateTime).Seconds())
		result = append(result, cdr)

	}

	return result, nil

}
