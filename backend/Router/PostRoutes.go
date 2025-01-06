package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaidanpoin/blog-go/Controller"
	"github.com/zaidanpoin/blog-go/Middleware"
)

func PostRoutes(router *gin.RouterGroup) {

	router.GET("/post", Controller.GetPosts)
	router.GET("/post/:id", Controller.GetPostByID)
	router.POST("/post", Middleware.CheckLogin(), Controller.CreatePost)
	router.DELETE("/post/:id", Middleware.UserMiddleware(), Controller.DeletePost)
	router.PATCH("/post/:id", Middleware.CheckLogin(), Controller.UpdatePost)

}
