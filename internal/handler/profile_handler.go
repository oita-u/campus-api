package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oita-u/campus-api/internal/service"
)

type ProfileHandler struct {
	profileService *service.ProfileService
}

func NewProfileHandler(profileService *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "認証に失敗しました。再度ログインしてください。",
		})
		return
	}

	profile, err := h.profileService.GetProfile(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}

