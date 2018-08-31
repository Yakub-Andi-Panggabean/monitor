package util

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type firstLogger struct {
	Filename string
}

var transactionLog = logrus.New()
var errorLog = logrus.New()

func init() {

	transactionalLogger := lumberjack.Logger{
		Filename:   GetConfig("logger.transactional").(string),
		MaxSize:    int(GetConfig("logger.maxSize").(int64)),
		MaxAge:     int(GetConfig("logger.MaxAge").(int64)),
		MaxBackups: int(GetConfig("logger.MaxBackups").(int64)),
	}

	errorLogger := lumberjack.Logger{
		Filename:   GetConfig("logger.error").(string),
		MaxSize:    int(GetConfig("logger.maxSize").(int64)),
		MaxAge:     int(GetConfig("logger.MaxAge").(int64)),
		MaxBackups: int(GetConfig("logger.MaxBackups").(int64)),
	}

	transactionLog.SetOutput(&transactionalLogger)
	errorLog.SetOutput(&errorLogger)

}

//NewFirstLogger create logger
func NewFirstLogger(f string) *firstLogger {

	return &firstLogger{
		Filename: f,
	}

}

//Debug logging with debug level
func (f *firstLogger) Debug(message interface{}) {
	transactionLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Debug(message)

}

//Info logging with Info level
func (f *firstLogger) Info(message interface{}) {
	transactionLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Info(message)
}

//Warn logging with Warn level
func (f *firstLogger) Warn(message interface{}) {
	transactionLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Warn(message)
}

//Error logging with Error level
func (f *firstLogger) Error(message interface{}) {

	transactionLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Error(message)

	errorLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Error(message)

}
