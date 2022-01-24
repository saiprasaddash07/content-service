package v1

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/content-service.git/config"
	"github.com/saiprasaddash07/content-service.git/constants"
	"github.com/saiprasaddash07/content-service.git/controllers/v1/contentServices"
	"github.com/saiprasaddash07/content-service.git/controllers/v1/utils"
	"github.com/saiprasaddash07/content-service.git/helpers/request"
	"github.com/saiprasaddash07/content-service.git/helpers/response"
	"github.com/saiprasaddash07/content-service.git/helpers/util"
)

func UploadCSVFile(c *gin.Context) {
	userId := c.Query("userId")
	file, header, err := c.Request.FormFile("csvFile")

	if !util.ValidateCSV(header.Filename) {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_CSV_REQUEST)))
		return
	}

	if header.Size > constants.MAX_FILE_SIZE {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.FILE_SIZE_EXCEEDED)))
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_CSV_REQUEST)))
		return
	}

	if err := contentServices.UploadFile(file, userId); err != nil {
		c.JSON(http.StatusInternalServerError, util.SendErrorResponse(err))
		return
	}

	res := response.Response{
		Status:  constants.API_SUCCESS_STATUS,
		Message: constants.UPLOAD_CSV_SUCCESS_MESSAGE,
	}
	c.JSON(http.StatusOK, util.StructToJSON(res))
}

func PostContentHandler(c *gin.Context) {
	contentFromContext, ok := c.Get("content")
	if !ok {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
		return
	}
	content := contentFromContext.(*request.Content)

	if err := utils.ValidateContentDetails(content, constants.API_TYPE_CREATE_CONTENT); err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}

	contentRes, err := contentServices.CreateContent(content)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}

	createResponse := &response.Content{
		ContentId: contentRes.ContentId,
		Title:     contentRes.Title,
		Story:     contentRes.Story,
		UserId:    contentRes.UserId,
	}

	res := response.Response{
		Status:  constants.API_SUCCESS_STATUS,
		Message: constants.CREATE_CONTENT_MESSAGE,
		Result:  createResponse,
	}
	c.JSON(http.StatusOK, util.StructToJSON(res))
}

func EditContentHandler(c *gin.Context) {
	contentFromContext, ok := c.Get("content")
	log.Println(contentFromContext)
	if !ok {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_REQUEST)))
		return
	}
	content := contentFromContext.(*request.Content)

	contentRes, err := contentServices.UpdateContent(content)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}

	createResponse := &response.Content{
		ContentId: contentRes.ContentId,
		Title:     contentRes.Title,
		Story:     contentRes.Story,
		UserId:    contentRes.UserId,
	}

	res := response.Response{
		Status:  constants.API_SUCCESS_STATUS,
		Message: constants.EDIT_CONTENT_MESSAGE,
		Result:  createResponse,
	}
	c.JSON(http.StatusOK, util.StructToJSON(res))
}

func DeleteContentHandler(c *gin.Context) {
	contentId, err := strconv.ParseInt(c.Query("contentId"), 10, 64)
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(errors.New(constants.INVALID_CONTENT_ID)))
		return
	}

	var content request.Content
	content.ContentId = contentId
	content.UserId = userId

	_, err = contentServices.DeleteContent(&content)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, util.SendSuccessResponse(constants.DELETE_CONTENT_MESSAGE))
}

func FetchNewContents(c *gin.Context) {
	contentRes, err := contentServices.FetchNewContents(config.Get().SizeOfNewContents)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, util.StructToJSON(contentRes))
}

func FetchTopContents(c *gin.Context) {
	contentRes, err := contentServices.FetchTopContents()
	if err != nil {
		c.JSON(http.StatusBadRequest, util.SendErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, util.StructToJSON(contentRes))
}