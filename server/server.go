package server

import (
	"github.com/saiprasaddash07/content-service.git/config"
	"github.com/saiprasaddash07/content-service.git/services/db"
	"github.com/saiprasaddash07/content-service.git/services/logger"
)

func Init() {
	config := config.Get()
	logger.InitLogger()
	db.Init()
	// redis.Init()
	r := NewRouter()
	r.Run(":" + config.ServerPort)
}
