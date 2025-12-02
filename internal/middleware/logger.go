package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/oita-u/campus-api/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// エラー（400以上）のみログ出力
		if status >= 400 {
			log := logger.Get().With(
				zap.Int("status", status),
				zap.String("method", method),
				zap.String("path", path),
				zap.Duration("latency", latency),
			)

			if errorMessage != "" {
				log = log.With(zap.String("error", errorMessage))
			}

			if status >= 500 {
				log.Error("HTTP Request")
			} else {
				log.Warn("HTTP Request")
			}
		}
	}
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				log := logger.Get().With(
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
				)

				if isBrokenPipe(err) {
					log.Error(c.Request.URL.Path)
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				log.Error("[Recovery from panic]", zap.String("stack", string(debug.Stack())))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func isBrokenPipe(err interface{}) bool {
	ne, ok := err.(*net.OpError)
	if !ok {
		return false
	}
	se, ok := ne.Err.(*os.SyscallError)
	if !ok {
		return false
	}
	errMsg := strings.ToLower(se.Error())
	return strings.Contains(errMsg, "broken pipe") || strings.Contains(errMsg, "connection reset by peer")
}
