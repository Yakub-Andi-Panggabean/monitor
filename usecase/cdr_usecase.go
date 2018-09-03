package usecase

import (
	"firstwap.com/sms-monitoring-api/entity"
	"firstwap.com/sms-monitoring-api/repository"
	"firstwap.com/sms-monitoring-api/util"
)

var log = util.NewFirstLogger("cdr_usecase")

type cdrUsecase struct {
	cdrRepo repository.CdrRepsitory
}

//NewCDRUsecase create cdr usecase instance
func NewCDRUsecase(r repository.CdrRepsitory) CDRUsecase {

	return &cdrUsecase{
		cdrRepo: r,
	}

}

func (r *cdrUsecase) FindDeliveryMetric() (*entity.CDR, error) {

	cdr := r.cdrRepo.FetchLatestUpdatedCDR()
	messageIDDateTime, _ := util.ParseQueueMessageID(cdr.QueueMessageID)
	time, err := util.ConvertDateStringToTime(messageIDDateTime)

	log.Info(cdr)

	if err != nil {
		return nil, err
	}

	return &entity.CDR{
		DateTime:          cdr.DateTime,
		QueueMessageID:    cdr.QueueMessageID,
		SubmmitedDateTime: cdr.SubmmitedDateTime,
		Interval:          int64(time.Sub(cdr.DateTime).Seconds()),
	}, nil

}
