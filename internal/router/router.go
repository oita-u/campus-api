package router

import (
	"github.com/oita-u/campus-api/internal/handler"
	"github.com/oita-u/campus-api/internal/middleware"
	"github.com/oita-u/campus-api/internal/repository"
	"github.com/oita-u/campus-api/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(middleware.GinLogger())
	r.Use(middleware.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)
	userHandler := handler.NewUserHandler(userService)

	profileRepo := repository.NewProfileRepository()
	profileService := service.NewProfileService(profileRepo)
	profileHandler := handler.NewProfileHandler(profileService)

	api := r.Group("/api")
	v1 := api.Group("/v1")

	authPublic := v1.Group("/auth")
	authPublic.POST("/login", authHandler.Login)
	authPublic.POST("/register", userHandler.Create)

	profile := v1.Group("/profile")
	profile.Use(middleware.JWT())
	{
		profile.GET("", profileHandler.GetProfile)
	}

	return r
}
