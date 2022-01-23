package contentServices

import (
	"encoding/csv"
	"errors"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/saiprasaddash07/content-service.git/constants"
	"github.com/saiprasaddash07/content-service.git/helpers/DAO"
	"github.com/saiprasaddash07/content-service.git/helpers/request"
)

func UploadFile(file multipart.File, userId string) error {
	reader := csv.NewReader(file)
	reqUserId, _ := strconv.ParseInt(string(userId), 10, 64)

	if ok := DAO.DoesUserExist(reqUserId); !ok {
		return errors.New(constants.ERROR_NO_USER_EXIST)
	}

	var stories []request.Content
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		if len(line) != 2 {
			// Every csv file should only contain 2 columns, i.e, title and description of story
			return errors.New(constants.INVALID_CSV_REQUEST)
		}
		stories = append(stories, request.Content{
			Title:  strings.TrimSpace(line[0]),
			Story:  strings.TrimSpace(line[1]),
			UserId: reqUserId,
		})
	}

	if err := DAO.SaveContent(stories); err != nil {
		return err
	}

	return nil
}

func CreateContent(content *request.Content) (*request.Content, error) {
	if err := DAO.CreateContent(content); err != nil {
		return nil, err
	}

	return content, nil
}

func UpdateContent(content *request.Content) (*request.Content, error) {
	if err := DAO.UpdateContent(content); err != nil {
		return nil, err
	}
	return content, nil
}