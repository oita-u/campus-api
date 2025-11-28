package router

import (
	"github.com/oita-u/campus-api/internal/handler"
	"github.com/oita-u/campus-api/internal/middleware"
	"github.com/oita-u/campus-api/internal/repository"
	"github.com/oita-u/campus-api/internal/service"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	studentRepo := &repository.StudentRepository{}
	studentService := &service.StudentService{Repo: studentRepo}
	studentHandler := &handler.StudentHandler{Service: studentService}

	v1 := r.Group("/v1")
	v1.GET("/ping", handler.Ping)
	v1.GET("/students/:id", studentHandler.GetByID)
	v1.GET("/students", studentHandler.List)

	// JWT保護ルート
	auth := v1.Group("/auth")
	auth.Use(middleware.JWT())
	{
		auth.GET("/me", handler.Me)
	}

	return r
}
