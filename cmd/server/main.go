package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/oita-u/campus-api/internal/config"
	"github.com/oita-u/campus-api/internal/logger"
	"github.com/oita-u/campus-api/internal/router"
)

func main() {
	config.Load()

	logger.Init()
	defer logger.Get().Sync()

	if config.C.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := router.New()

	addr := ":" + config.C.Port
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
