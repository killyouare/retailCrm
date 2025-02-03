package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

type logger struct {
	logger *zap.Logger
}

func (l logger) Debug(msg string, args ...interface{}) {
	l.logger.Log(zap.DebugLevel, fmt.Sprintf(msg, args...))
}

func (l logger) Info(msg string, args ...interface{}) {
	l.logger.Log(zap.InfoLevel, fmt.Sprintf(msg, args...))
}

func (l logger) Warn(msg string, args ...interface{}) {
	l.logger.Log(zap.WarnLevel, fmt.Sprintf(msg, args...))
}

func (l logger) Error(msg string, args ...interface{}) {
	l.logger.Log(zap.ErrorLevel, fmt.Sprintf(msg, args...))
}

func (l logger) Fatal(msg string, args ...interface{}) {
	l.logger.Log(zap.FatalLevel, fmt.Sprintf(msg, args...))
}

func New() (Logger, error) {
	development, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &logger{
		logger: development.WithOptions(zap.AddCallerSkip(1)),
	}, nil
}
