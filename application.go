package main

import (
	"firstwap.com/sms-monitoring-api/util"
)

var log = util.NewFirstLogger("main")

func main() {

	log.Error(util.GetConfig("logger.error"))

}
