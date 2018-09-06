package usecase

import (
	"firstwap.com/sms-monitoring-api/entity"
	"firstwap.com/sms-monitoring-api/repository"
	"firstwap.com/sms-monitoring-api/util"
)

var log = util.NewFirstLogger("cdr_usecase")

type cdrUsecase struct {
	cdrRepo repository.CdrRepository
}

//NewCDRUsecase create cdr usecase instance
func NewCDRUsecase(r repository.CdrRepository) CDRUsecase {

	return &cdrUsecase{
		cdrRepo: r,
	}

}

func (r *cdrUsecase) FindDeliveryMetric() ([]*entity.CDR, error) {

	result := make([]*entity.CDR, 0)
	cdrs := r.cdrRepo.FetchLatestUpdatedCDR()

	for _, cdr := range cdrs {

		messageIDDateTime, _ := util.ParseQueueMessageID(cdr.QueueMessageID)
		time, err := util.ConvertDateStringToTime(messageIDDateTime)

		if err != nil {
			return nil, err
		}

		cdr.Interval = int64(time.Sub(cdr.DateTime).Seconds())
		result = append(result, cdr)

	}

	return result, nil

}
