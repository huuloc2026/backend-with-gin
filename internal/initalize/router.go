package initalize

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huuloc2026/go-backend/internal/handler"
	"github.com/huuloc2026/go-backend/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("- AA executed")
		c.Next()
		fmt.Println("AA executed")
	}
}
func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("- BB executed")
		c.Next()
		fmt.Println("BB executed")
	}
}

func CC() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("- CC executed")
		c.Next()
		fmt.Println("CC executed")
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenticateMiddleware())
	// r.Use(AA(), BB(), CC())
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping:name", Pong)
		v1.GET("/users/", handler.NewUserHandler().GetUserByID)
	}
	return r

}

func Pong(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + "ssshello" + name,
	})
}
