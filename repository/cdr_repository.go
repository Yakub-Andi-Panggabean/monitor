package repository

import (
	"database/sql"
	"time"

	"firstwap.com/sms-monitoring-api/entity"
	"firstwap.com/sms-monitoring-api/util"
)

var log = util.NewFirstLogger("cdr_repository")

type mySqlCdrRepository struct {
	Db       sql.DB
	Provider string
}

//FetchLatestUpdatedCDR return slice of entity based on the number of databases which is scanned
func (m *mySqlCdrRepository) FetchLatestUpdatedCDR() []*entity.CDR {

	result := make([]*entity.CDR, 0)

	var cdr entity.CDR

	row := m.Db.QueryRow("Select SubmitDateTime,DateTime,QueueMessageID,DeliveryStatus FROM " + time.Now().Format("200601") + " where DeliveryStatus is not null order by ID desc limit 1")

	err := row.Scan(&cdr.SubmmitedDateTime, &cdr.DateTime, &cdr.QueueMessageID, &cdr.DeliveryCode)

	if err != nil {
		log.Error("an error occured ", err.Error())
	} else {

	}

	return result

}

//create new Instance of Mysql repository
func NewMySQLRepository(db sql.DB, provider string) CdrRepository {

	return &mySqlCdrRepository{
		Db:       db,
		Provider: provider,
	}

}
