package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saiprasaddash07/content-service.git/constants"
	"github.com/saiprasaddash07/content-service.git/controllers/v1/utils"
)

func GetRequestBodyContent(apiType string, contentRequiredFields []string, contentOptionalFields []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		var requestObj interface{}

		if err := context.ShouldBind(&requestObj); err != nil {
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": constants.INVALID_REQUEST,
			})
			return
		}

		contentJSON := requestObj.(map[string]interface{})
		content, ok := utils.ValidateAndParseContentFields(contentJSON, contentRequiredFields, contentOptionalFields)
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": constants.INVALID_REQUEST,
			})
			return
		}

		if apiType == constants.API_TYPE_EDIT_CONTENT {
			contentId, err := strconv.ParseInt(context.Query("contentId"), 10, 64)
			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"status":  constants.API_FAILED_STATUS,
					"message": constants.INVALID_CONTENT_ID,
				})
				return
			}
			content.ContentId = contentId
		}

		if err := utils.ValidateContentDetails(content, apiType); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  constants.API_FAILED_STATUS,
				"message": err.Error(),
			})
			return
		}

		context.Set("content", content)
		context.Next()
	}
}
