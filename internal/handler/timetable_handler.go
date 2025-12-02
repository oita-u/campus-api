package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oita-u/campus-api/internal/service"
)

type TimetableHandler struct {
	timetableService *service.TimetableService
}

func NewTimetableHandler(timetableService *service.TimetableService) *TimetableHandler {
	return &TimetableHandler{
		timetableService: timetableService,
	}
}

func (h *TimetableHandler) GetTimetable(c *gin.Context) {
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
	if yearStr == "" {
		year = time.Now().Year()
	} else {
		parsedYear, err := strconv.Atoi(yearStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "year は数値で指定してください。",
			})
			return
		}
		year = parsedYear
	}

	items, err := h.timetableService.GetTimetable(userID.(string), year, semester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "時間割の取得に失敗しました。",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"year":     year,
		"semester": semester,
		"items":    items,
	})
}
