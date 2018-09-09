package parser

import (
	"errors"
	"time"
)

const (
	//DateLayout constant for parsing message id date time
	DateLayout = "2006-01-02 15:04:05.000"
)

//ParseQueueMessageID parsing queueMessageID to extract SMS API datetime,
//queueMessageID format : 5GPI2018-08-17 02:14:59.489.fB3GN,PJA
func ParseQueueMessageID(queueMessageID string) (string, error) {

	//substring start from 5th character to 28th character
	if len(queueMessageID) < len("5GPI2018-08-17 02:14:59.489.fB3GN,PJA") {

		return "", errors.New("invalid queue message id format")

	}

	return queueMessageID[4:27], nil

}

//ConvertDateStringToTime parsing date time extracted from queue message id,
//date time format yyyy-MM-dd HH:mm:ss.SSS
func ConvertDateStringToTime(messageIDTime string) (time.Time, error) {

	return time.Parse(DateLayout, messageIDTime)
}


func Contains(arr []string,val string) bool{

	for _,existingValue:=range arr{

		if(val==existingValue){
			return true
		}

	}

	return false

}
