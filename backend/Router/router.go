package Router

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServeApps() {
	router := gin.Default()

	router.Use(cors.Default())
	router.Static("/uploads", "./uploads")
	apiRoutes := router.Group("/api/v1")
	{

		PostRoutes(apiRoutes)

		auth := apiRoutes.Group("/auth")
		{
			AuthRoutes(auth)
		}
	}

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}
