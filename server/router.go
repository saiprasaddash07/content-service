package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/saiprasaddash07/content-service.git/controllers/v1"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	version1 := router.Group("api/v1")
	{
		contentGroupV1 := version1.Group("content")
		{
			contentGroupV1.POST("/upload", v1.UploadCSVFile)
		}
	}

	return router
}
