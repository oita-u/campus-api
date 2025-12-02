package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oita-u/campus-api/internal/service"
)

type GradeHandler struct {
	gradeService *service.GradeService
}

func NewGradeHandler(gradeService *service.GradeService) *GradeHandler {
	return &GradeHandler{
		gradeService: gradeService,
	}
}

// GetGrades はログインユーザーの成績一覧を取得する
// GET /api/v1/grades?year=2024&semester=1st など
func (h *GradeHandler) GetGrades(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "認証に失敗しました。再度ログインしてください。",
		})
		return
	}

	yearStr := c.Query("year")
	semester := c.Query("semester")

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

	resp, err := h.gradeService.GetGrades(userID.(string), year, semester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "成績の取得に失敗しました。",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}


