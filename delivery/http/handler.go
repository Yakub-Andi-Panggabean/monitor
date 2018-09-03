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

	result, err := m.CdrUsecase.CalculateDeliveryInterval()

	if err != nil {

	}

	json, err := json.Marshal(Metric{

		Interval: int64(result),
	})

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

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
