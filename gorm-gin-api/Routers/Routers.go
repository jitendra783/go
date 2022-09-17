package Routers

import (
	"example/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/user")
	{
		v1.GET("/", Controllers.ListUser)
		v1.POST("/", Controllers.AddNewUser)
		v1.GET("/:id", Controllers.GetOneUser)
		v1.PUT("/:id", Controllers.PutOneUser)
		v1.DELETE("/:id", Controllers.DeleteUser)
	}

	return r
}
