package mysql

import (
	"database/sql"
	"time"

	"firstwap.com/sms-monitoring-api/application/domain"
	"firstwap.com/sms-monitoring-api/application/logger"
)

var log = logger.New("cdr_repository")

type MySQLStorage struct {
	Db       sql.DB
	Provider string
}

//FetchLatestUpdatedCDR return slice of entity based on the number of databases which is scanned
func (m *MySQLStorage) FetchLatestUpdatedCDR() []*domain.CDR {

	result := make([]*domain.CDR, 0)

	var cdr domain.CDR

	row := m.Db.QueryRow("Select SubmitDateTime,DateTime,QueueMessageID,DeliveryStatus FROM " + time.Now().Format("200601") + " where DeliveryStatus is not null order by ID desc limit 1")

	err := row.Scan(&cdr.SubmmitedDateTime, &cdr.DateTime, &cdr.QueueMessageID, &cdr.DeliveryCode)

	if err != nil {
		log.Error("an error occured ", err.Error())
	} else {

	}

	return result

}

//create new Instance of Mysql repository
func New(db sql.DB, provider string) *MySQLStorage {

	return &MySQLStorage{
		Db:       db,
		Provider: provider,
	}

}
