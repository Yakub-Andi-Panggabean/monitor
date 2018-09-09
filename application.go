package main

import (
	"net/http"
	"strconv"

	"firstwap.com/sms-monitoring-api/application/config"
	monitorHttp "firstwap.com/sms-monitoring-api/application/http"
	"firstwap.com/sms-monitoring-api/application/logger"

	"github.com/gorilla/mux"
)

var log = logger.New("main")

func main() {

	serverPort := ":" + strconv.Itoa(int(config.Get("server.port").(int64)))

	r := mux.NewRouter()

	monitorHttp.New(r)

	err := http.ListenAndServe(serverPort, r)

	if err != nil {

		log.Error(config.Get("logger.error"))

	} else {

		log.Info("Server serve port ", serverPort)

	}

}
