package logger

import "go.uber.org/zap"

var L *zap.Logger

func Init() {
	l, _ := zap.NewProduction()
	L = l
}

func Get() *zap.Logger {
	return L
}
