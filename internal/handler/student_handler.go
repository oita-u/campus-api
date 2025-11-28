package handler

import (
	"net/http"
	"strconv"

	"github.com/oita-u/campus-api/internal/service"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	Service *service.StudentService
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	s, err := h.Service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

func (h *StudentHandler) List(c *gin.Context) {
	list, err := h.Service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	c.JSON(http.StatusOK, list)
}
