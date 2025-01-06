package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaidanpoin/blog-go/Controller"
)

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", Controller.Register)

	router.POST("/login", Controller.Login)
}
