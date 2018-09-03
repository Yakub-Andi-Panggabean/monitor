package main

import (
	"net/http"
	"strconv"

	"firstwap.com/sms-monitoring-api/util"
	"github.com/gorilla/mux"
)

var log = util.NewFirstLogger("main")

func main() {

	serverPort := ":" + strconv.Itoa(int(util.GetConfig("server.port").(int64)))

	r := mux.NewRouter()
	//delivery.NewHTTPMetricHandler(r)

	err := http.ListenAndServe(serverPort, r)

	if err != nil {
		log.Error(util.GetConfig("logger.error"))
	} else {
		log.Info("Server serve port ", serverPort)
	}

}
