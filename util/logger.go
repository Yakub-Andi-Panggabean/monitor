package util

import (
	"time"

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
		Filename:   "d:\\var\\log\\monitoring." + time.Now().Local().Format("2006-01-02") + ".log",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 7,
	}

	errorLogger := lumberjack.Logger{
		Filename:   "d:\\var\\log\\monitoring.err." + time.Now().Local().Format("2006-01-02") + ".log",
		MaxSize:    100,
		MaxAge:     30,
		MaxBackups: 7,
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
