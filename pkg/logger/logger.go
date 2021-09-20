package logger

import (
	"github.com/ZhansultanS/myLMS/backend/internal/config"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
}

func NewLogrus(env string) Logger {
	l := logrus.New()

	if env == config.EnvProduction {
		l.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: true,
		})
	}

	return l
}
