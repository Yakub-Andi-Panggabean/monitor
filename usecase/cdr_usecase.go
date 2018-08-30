package usecase

import (
	"firstwap.com/sms-monitoring-api/entity"
	"firstwap.com/sms-monitoring-api/repository"
	"firstwap.com/sms-monitoring-api/util"
	log "github.com/sirupsen/logrus"
)

type cdrUsecase struct {
	cdrRepo repository.CdrRepsitory
}

//NewCDRUsecase create cdr usecase instance
func NewCDRUsecase(r repository.CdrRepsitory) *cdrUsecase {

	return &cdrUsecase{
		cdrRepo: r,
	}

}

func (r *cdrUsecase) calculateDeliveryInterval(cdr entity.CDR) (float64, error) {

	messageIDDateTime := util.ParseQueueMessageID(cdr.QueueMessageID)
	time, err := util.ConvertDateStringToTime(messageIDDateTime)

	log.Info(cdr)

	if err != nil {
		return 0, err
	} else {
		return time.Sub(cdr.DateTime).Seconds(), nil
	}

}
