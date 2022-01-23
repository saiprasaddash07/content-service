package DAO

import (
	"fmt"
	"log"

	"github.com/saiprasaddash07/content-service.git/constants"
	"github.com/saiprasaddash07/content-service.git/services/db"
)

func WriteBatch(valuesToWrite []interface{}, writeSql string) error {
	writeSql = writeSql[0 : len(writeSql)-2]
	stmt, err := db.GetClient(constants.DB_WRITER).Prepare(writeSql)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	writeSql += fmt.Sprintf(" ON DUPLICATE KEY UPDATE contentId = contentId")
	_, err = stmt.Exec(valuesToWrite...)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
