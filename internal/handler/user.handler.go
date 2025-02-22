package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/huuloc2026/go-backend/internal/service"
	"github.com/huuloc2026/go-backend/pkg/response"
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
	fmt.Println("- UserHandler executed")
	response.SuccessResponse(c, 2000, []string{"hello", "loc"})

}
