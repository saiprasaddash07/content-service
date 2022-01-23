package utils

import (
	"errors"
	"fmt"

	"github.com/saiprasaddash07/content-service.git/constants"
	"github.com/saiprasaddash07/content-service.git/helpers/DAO"
	"github.com/saiprasaddash07/content-service.git/helpers/request"
	"github.com/saiprasaddash07/content-service.git/helpers/util"
)

func ValidateAndParseContentFields(contentJSON map[string]interface{}, requiredFields []string, optionalFields []string) (*request.Content, bool) {
	lengthDiffRequiredFieldsAndcontentJSON := len(contentJSON) - len(requiredFields)
	if lengthDiffRequiredFieldsAndcontentJSON < 0 || len(optionalFields) < lengthDiffRequiredFieldsAndcontentJSON {
		return nil, false
	}

	countOfReqFields := len(requiredFields)
	var content request.Content
	for k, v := range contentJSON {
		if util.Contains(requiredFields, k) {
			countOfReqFields--
		} else if !util.Contains(optionalFields, k) {
			return nil, false
		}

		valueType := fmt.Sprintf("%T", v)
		switch k {
		case "title":
			if valueType == "string" {
				content.Title = v.(string)
			} else {
				return &content, false
			}
		case "story":
			if valueType == "string" {
				content.Story = v.(string)
			} else {
				return &content, false
			}
		case "userId":
			if valueType == "float64" && util.IsInteger(v.(float64)) {
				content.UserId = int64(v.(float64))
			} else {
				return &content, false
			}
		default:
			return nil, false
		}
	}
	if countOfReqFields == 0 {
		return &content, true
	}
	return nil, false
}

func ValidateContentDetails(content *request.Content) error {
	if len(content.Title) < constants.MIN_LENGTH_OF_CONTENT_TITLE || len(content.Title) > constants.MAX_LENGTH_OF_CONTENT_TITLE {
		return errors.New(constants.INVALID_REQUEST)
	}
	if len(content.Story) == 0 {
		return errors.New(constants.INVALID_REQUEST)
	}
	if ok := DAO.DoesUserExist(content.UserId); !ok {
		return errors.New(constants.INVALID_USER_ID)
	}
	return nil
}
