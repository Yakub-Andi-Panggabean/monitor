package application

import "firstwap.com/sms-monitoring-api/application/domain"

type CdrStorageFetcher interface {
	FetchLatestUpdatedCDR() []*domain.CDR
}
