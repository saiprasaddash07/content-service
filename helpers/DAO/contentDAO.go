package DAO

import (
	"log"
	"strings"

	"github.com/saiprasaddash07/content-service.git/constants"
	"github.com/saiprasaddash07/content-service.git/helpers/request"
	"github.com/saiprasaddash07/content-service.git/services/db"
)

func SaveContent(stories []request.Content) error {
	query := "INSERT INTO content (title, story, userId) VALUES "
	valuesToWrite := []interface{}{}
	for i := 0; i < len(stories); i++ {
		query += "(?, ?, ?), "
		valuesToWrite = append(valuesToWrite, stories[i].Title, stories[i].Story, stories[i].UserId)
	}

	err := WriteBatch(valuesToWrite, query)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func DoesUserExist(userId int64) bool {
	var count int64
	err := db.GetClient(constants.DB_READER).QueryRow("SELECT COUNT(*) AS count FROM users WHERE userId=?;", userId).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func CreateContent(content *request.Content) error {
	content.Title = strings.TrimSpace(content.Title)
	content.Story = strings.TrimSpace(content.Story)
	rows, err := db.GetClient(constants.DB_WRITER).Exec("INSERT INTO content (title, story, userId) VALUES (?,?,?);", content.Title, content.Story, content.UserId)
	if err != nil {
		return err
	}
	contentId, err := rows.LastInsertId()
	if err != nil {
		return err
	}
	content.ContentId = contentId
	return nil
}