package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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

	if err := utils.ValidateContentDetails(content); err != nil {
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
