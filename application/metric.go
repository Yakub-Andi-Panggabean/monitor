package application

import "firstwap.com/sms-monitoring-api/application/domain"

//DeliveryMetrciFinder is an interface used to find delivery metrics from datasource.
type DeliveryMetricFinder interface {
	FindDeliveryMetric() ([]*domain.CDR, error)
}
