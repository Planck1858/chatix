package logging

import (
	"go.uber.org/zap"
)

var l *zap.SugaredLogger

type Logger struct {
	*zap.SugaredLogger
}

func Init() {
	zapLog, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	l = zapLog.Sugar()
}

func GetLogger() Logger {
	return Logger{l}
}
