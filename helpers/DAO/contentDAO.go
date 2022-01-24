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

func DoesContentBelongsToUser(userId int64, contentId int64) bool {
	var count int64
	err := db.GetClient(constants.DB_READER).QueryRow("SELECT COUNT(*) AS count FROM content WHERE userId=? AND contentId=?;", userId, contentId).Scan(&count)
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

func UpdateContent(content *request.Content) error {
	_, err := db.GetClient(constants.DB_WRITER).Exec("UPDATE content SET title=?, story=? WHERE contentId = ? AND userId=?", content.Title, content.Story, content.ContentId, content.UserId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteContent(content *request.Content) error {
	_, err := db.GetClient(constants.DB_WRITER).Exec("UPDATE content SET isDeleted= ? WHERE contentId=? AND userId=?;", "true", content.ContentId, content.UserId)
	if err != nil {
		return err
	}
	return nil
}

func FetchNewContents(size int) ([]request.Content, error) {
	var contents []request.Content
	rows, err := db.GetClient(constants.DB_READER).Query("SELECT contentId, title, story, userId FROM content WHERE isDeleted=? ORDER BY contentId DESC LIMIT ?;", "false", size)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var content request.Content
		err := rows.Scan(&content.ContentId, &content.Title, &content.Story, &content.UserId)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}

func FetchContents(ids []interface{}) ([]request.Content, error) {
	var contents []request.Content
	query := `SELECT contentId, title, story, userId FROM content WHERE contentId IN (?` + strings.Repeat(",?", len(ids)-1) + ");"
	log.Println(query)
	rows, err := db.GetClient(constants.DB_READER).Query(query, ids...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var content request.Content
		err := rows.Scan(&content.ContentId, &content.Title, &content.Story, &content.UserId)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}