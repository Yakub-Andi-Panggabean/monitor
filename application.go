package main

import (
	"firstwap.com/sms-monitoring-api/util"
)

var logger = *util.NewFirstLogger("main")

func main() {

	//fmt.Println(util.ParseQueueMessageID("5GPI2018-08-17 02:14:59.489.fB3GN,PJA"))
	res, err := util.ParseQueueMessageID("")
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info(res)
	}

}
