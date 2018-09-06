package main

import (
	"net/http"
	"strconv"

	delivery "firstwap.com/sms-monitoring-api/delivery/http"
	"firstwap.com/sms-monitoring-api/repository"
	"firstwap.com/sms-monitoring-api/util"
	"github.com/gorilla/mux"

	"firstwap.com/sms-monitoring-api/usecase"
)

var log = util.NewFirstLogger("main")
var cdrRepository = repository.NewMySQLRepository()

func main() {

	serverPort := ":" + strconv.Itoa(int(util.GetConfig("server.port").(int64)))

	r := mux.NewRouter()

	cdrUsecase := usecase.NewCDRUsecase(cdrRepository)

	delivery.NewHTTPMetricHandler(r, cdrUsecase)

	err := http.ListenAndServe(serverPort, r)

	if err != nil {

		log.Error(util.GetConfig("logger.error"))

	} else {

		log.Info("Server serve port ", serverPort)

	}

}
