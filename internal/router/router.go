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

	gradeRepo := repository.NewGradeRepository()
	gradeService := service.NewGradeService(gradeRepo)
	gradeHandler := handler.NewGradeHandler(gradeService)

	timetableRepo := repository.NewTimetableRepository()
	timetableService := service.NewTimetableService(timetableRepo)
	timetableHandler := handler.NewTimetableHandler(timetableService)

	courseRepo := repository.NewCourseRepository()
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

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

	grades := v1.Group("/grades")
	grades.Use(middleware.JWT())
	{
		grades.GET("", gradeHandler.GetGrades)
	}

	timetable := v1.Group("/timetable")
	timetable.Use(middleware.JWT())
	{
		timetable.GET("", timetableHandler.GetTimetable)
	}

	courses := v1.Group("/courses")
	courses.Use(middleware.JWT())
	{
		courses.GET("", courseHandler.SearchCourses)
	}

	return r
}
