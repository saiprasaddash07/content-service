package server

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	version1 := router.Group("api/v1")
	{
		userGroupV1 := version1.Group("user")
		{
			userGroupV1.POST("/register")
		}
	}

	return router
}
