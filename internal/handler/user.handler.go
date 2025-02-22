package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huuloc2026/go-backend/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// uc -> user controller
func (uc *UserHandler) GetUserByID(c *gin.Context) {
	// id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInforUserService(),
	})

}
