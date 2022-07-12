package logrus

import (
	"github.com/sirupsen/logrus"
	"go_web_boilerplate/internal/pkg/logger"
	"io"
)

type log struct {
	logrus logrus.Logger
}

func New(writer io.Writer, level logrus.Level) logger.Logger {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(writer)
	logrus.SetLevel(level)
	return log{
		logrus: *logrus.New(),
	}
}

func (l log) Error(msg string, kv ...interface{}) {
	l.logrus.Error(msg, kv)
}

func (l log) Warn(msg string, kv ...interface{}) {
	l.logrus.Warn(msg, kv)
}

func (l log) Info(msg string, kv ...interface{}) {
	l.logrus.Info(msg, kv)

}

func (l log) Sync() {

}
