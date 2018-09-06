package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"firstwap.com/sms-monitoring-api/usecase"
	"github.com/gorilla/mux"
)

//MetricsHandler http metric handler struct
type MetricsHandler struct {
	CdrUsecase usecase.CDRUsecase
}

//MetricsHandler is the handler which will be used for handler metric request via http protocol
func (m *MetricsHandler) MetricsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	metrics := make([]Metric, 0)
	result, err := m.CdrUsecase.FindDeliveryMetric()

	if err != nil {

		w.Write(ErrorHandler(0, err.Error()))

	} else {

		for i := 0; len(result); i++ {

			metric := Metric{
				Interval:          int64(result[i].Interval),
				AcceptedDateTime:  result[i].SubmmitedDateTime,
				DeliveredDateTime: result[i].DateTime,
				Provider:          "",
				SmsAPIMessageID:   result[i].QueueMessageID,
			}

			metrics = append(metrics, metric)

		}

		w.Write(metrics)

	}

}

//HomeHandler is the handler which will be executed for home path request
func (m *MetricsHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<h1>Sms Monitoring Api</h1>")

}

//NewHTTPMetricHandler initiate metric http handler
func NewHTTPMetricHandler(r *mux.Router, u usecase.CDRUsecase) {

	handler := &MetricsHandler{
		CdrUsecase: u,
	}

	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/metrics", handler.MetricsHandler)

}

func ErrorHandler(errorCode int, errorMessage string) []byte {

	res, _ := json.Marshal(ErrorResponse{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	})

	return res

}
