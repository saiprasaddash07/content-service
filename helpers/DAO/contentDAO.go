package DAO

import (
	"log"

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
