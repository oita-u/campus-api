package handler

import (
	"net/http"

	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/oita-u/campus-api/internal/auth"
	"github.com/oita-u/campus-api/internal/logger"
	"github.com/oita-u/campus-api/internal/service"
	"go.uber.org/zap"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string    `json:"token"`
	User  *userInfo `json:"user"`
}

type userInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "リクエストパラメータが正しくありません",
		})
		return
	}

	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "メールアドレスの形式が正しくありません",
		})
		return
	}

	if len(req.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "パスワードは8文字以上である必要があります",
		})
		return
	}

	user, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		logger.Get().Error("ログイン失敗",
			zap.String("email", req.Email),
			zap.Error(err),
		)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "トークンの生成に失敗しました",
		})
		return
	}

	c.JSON(http.StatusOK, loginResponse{
		Token: token,
		User: &userInfo{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	})
}
