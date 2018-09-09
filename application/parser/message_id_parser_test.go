package parser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInvalidArgumentParseQueueMessageID(t *testing.T) {

	_, err := ParseQueueMessageID("5GPI2018-08-17 02:14:59")

	assert.NotNil(t, err)

	_, errEmpty := ParseQueueMessageID("")

	assert.NotNil(t, errEmpty)

}

func TestParseQueueMessageID(t *testing.T) {
	idTime, _ := ParseQueueMessageID("5GPI2018-08-17 02:14:59.489.fB3GN,PJA")
	assert.Equal(t, "2018-08-17 02:14:59.489", idTime)
}

func TestConvertDateStringToTime(t *testing.T) {

	time, _ := time.Parse(DateLayout, "2018-08-17 02:14:59.489")
	convertedTime, _ := ConvertDateStringToTime("2018-08-17 02:14:59.489")
	assert.Equal(t, time, convertedTime)

}
