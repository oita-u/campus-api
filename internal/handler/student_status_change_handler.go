package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oita-u/campus-api/internal/service"
)

type StudentStatusChangeHandler struct {
	Service *service.StudentStatusChangeService
}

func (h *StudentStatusChangeHandler) GetByStudentNumber(c *gin.Context) {
	studentNumber := c.Param("id")

	changes, err := h.Service.GetHistory(studentNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, changes)
}
