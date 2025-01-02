package Router

import (
	"fmt"

	"github.com/zaidanpoin/blog-go/Controller"
	"github.com/zaidanpoin/blog-go/Middleware"

	"github.com/gin-gonic/gin"
)

func ServeApps() {
	router := gin.Default()

	authRoutes := router.Group("/auth")
	{
		AuthRoutes(authRoutes)
	}

	router.GET("/user", Middleware.UserMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", Controller.Register)

	router.POST("/login", Controller.Login)
}
