package server

import (
	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/content-service.git/constants"
	v1 "github.com/saiprasaddash07/content-service.git/controllers/v1"
	"github.com/saiprasaddash07/content-service.git/controllers/v1/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	version1 := router.Group("api/v1")
	{
		contentGroupV1 := version1.Group("content")
		{
			contentGroupV1.POST("/upload", v1.UploadCSVFile)
			contentGroupV1.POST("/create", middlewares.GetRequestBodyContent(constants.API_TYPE_CREATE_CONTENT, constants.CREATE_CONTENT_REQUIRED_FIELDS, constants.CREATE_CONTENT_OPTIONAL_FIELDS), v1.PostContentHandler)
			contentGroupV1.POST("/edit", middlewares.GetRequestBodyContent(constants.API_TYPE_EDIT_CONTENT, constants.EDIT_CONTENT_REQUIRED_FIELDS, constants.EDIT_CONTENT_OPTIONAL_FIELDS), v1.EditContentHandler)
			contentGroupV1.POST("/delete", v1.DeleteContentHandler)
			contentGroupV1.GET("/new/contents", v1.FetchNewContents)
			contentGroupV1.GET("/top/contents", v1.FetchTopContents)
		}
	}

	return router
}
