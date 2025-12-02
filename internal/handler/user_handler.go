package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oita-u/campus-api/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

type createUserReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func (h *UserHandler) Create(c *gin.Context) {
	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "リクエストパラメータが正しくありません"})
		return
	}

	u, err := h.svc.Create(req.Email, req.Password, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) Me(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証に失敗しました。再度ログインしてください。"})
		return
	}

	u, err := h.svc.Get(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "ユーザーが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, u)
}
