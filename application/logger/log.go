package logger

import (
	"firstwap.com/sms-monitoring-api/application/config"
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
		Filename:   config.Get("logger.transactional").(string),
		MaxSize:    int(config.Get("logger.maxSize").(int64)),
		MaxAge:     int(config.Get("logger.MaxAge").(int64)),
		MaxBackups: int(config.Get("logger.MaxBackups").(int64)),
	}

	errorLogger := lumberjack.Logger{
		Filename:   config.Get("logger.error").(string),
		MaxSize:    int(config.Get("logger.maxSize").(int64)),
		MaxAge:     int(config.Get("logger.MaxAge").(int64)),
		MaxBackups: int(config.Get("logger.MaxBackups").(int64)),
	}

	transactionLog.SetOutput(&transactionalLogger)
	errorLog.SetOutput(&errorLogger)

}

//NewFirstLogger create logger
func New(f string) *firstLogger {

	return &firstLogger{
		Filename: f,
	}

}

//Debug logging with debug level
func (f *firstLogger) Debug(message ...interface{}) {
	transactionLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Debug(message)

}

//Info logging with Info level
func (f *firstLogger) Info(message ...interface{}) {
	transactionLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Info(message)
}

//Warn logging with Warn level
func (f *firstLogger) Warn(message ...interface{}) {
	transactionLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Warn(message)
}

//Error logging with Error level
func (f *firstLogger) Error(message ...interface{}) {

	transactionLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Error(message)

	errorLog.WithFields(logrus.Fields{
		"file": f.Filename,
	}).Error(message)

}
