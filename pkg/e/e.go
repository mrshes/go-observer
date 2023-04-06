package e

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	log *logrus.Logger
)

func Log() *logrus.Logger {
	return log
}

func New() *logrus.Logger {
	fmt.Println("Init logger!")
	logLevel := os.Getenv("LOGGER_LEVEL")
	logger := logrus.New()
	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	logger.SetLevel(getLevelForLogrus(logLevel))
	logger.Info("Init logger!")
	log = logger
	return log
}

func getLevelForLogrus(log_level string) logrus.Level {
	switch strings.ToUpper(log_level) {
	case "debug":
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}

func WithDataInfo(msg string, data interface{}) {
	Log().WithFields(logrus.Fields{
		"data": data,
	}).Info(msg)
}

func Error(msg string, err error) {
	Log().Error(msg+": ", err)
}

func Info(args ...interface{}) {
	Log().Info(args...)
}
