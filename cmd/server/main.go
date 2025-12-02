package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/oita-u/campus-api/internal/config"
	"github.com/oita-u/campus-api/internal/db"
	"github.com/oita-u/campus-api/internal/logger"
	"github.com/oita-u/campus-api/internal/router"
)

func main() {
	config.Load()
	db.Init()

	logger.Init()
	defer logger.Get().Sync()

	addr := ":" + config.C.Port

	if config.C.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := router.New()

	logger.Get().Info("サーバーを起動しています", zap.String("address", addr))
	if err := r.Run(addr); err != nil {
		logger.Get().Fatal("サーバーの起動に失敗しました", zap.Error(err))
	}
}
