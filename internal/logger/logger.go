package logger

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.Logger

func Init() {
	config := zap.NewDevelopmentConfig()
	L, _ = config.Build()
	L = L.WithOptions(zap.AddCaller(), zap.AddStacktrace(zap.DebugLevel))

	config.DisableCaller = true
	config.DisableStacktrace = true
	ginLogger, _ := config.Build()

	gin.DefaultWriter = &zapWriter{logger: ginLogger, level: zapcore.InfoLevel}
	gin.DefaultErrorWriter = &zapWriter{logger: ginLogger, level: zapcore.ErrorLevel}
}

func Get() *zap.Logger {
	return L
}

type zapWriter struct {
	logger *zap.Logger
	level  zapcore.Level
}

func (w *zapWriter) Write(p []byte) (n int, err error) {
	msg := strings.TrimSuffix(string(p), "\n")
	if w.level == zapcore.ErrorLevel {
		w.logger.Error(msg)
	} else {
		w.logger.Info(msg)
	}
	return len(p), nil
}
