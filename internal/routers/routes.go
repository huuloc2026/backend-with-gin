package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huuloc2026/go-backend/internal/handler"
)

func NewRoute() *gin.Engine {
	r := gin.Default()
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
