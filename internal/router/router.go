package router

import (
	"github.com/oita-u/campus-api/internal/handler"
	"github.com/oita-u/campus-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.GET("/ping", handler.Ping)

	// JWT保護ルート
	auth := api.Group("/auth")
	auth.Use(middleware.JWT())
	{
		auth.GET("/me", handler.Me)
	}

	return r
}
