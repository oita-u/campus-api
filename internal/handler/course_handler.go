package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oita-u/campus-api/internal/service"
)

type CourseHandler struct {
	courseService *service.CourseService
}

func NewCourseHandler(courseService *service.CourseService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
	}
}

func (h *CourseHandler) SearchCourses(c *gin.Context) {
	yearStr := c.Query("year")
	semester := c.Query("semester")
	keyword := c.Query("q")

	var year int
	if yearStr != "" {
		parsedYear, err := strconv.Atoi(yearStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "year は数値で指定してください。",
			})
			return
		}
		year = parsedYear
	}

	courses, resolvedYear, resolvedSemester, err := h.courseService.SearchCourses(year, semester, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "シラバス一覧の取得に失敗しました。",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"year":     resolvedYear,
		"semester": resolvedSemester,
		"courses":  courses,
	})
}
