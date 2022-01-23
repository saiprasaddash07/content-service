package logger

import (
	"sync"

	"github.com/saiprasaddash07/content-service.git/config"
	"go.uber.org/zap"
)

var logger *zap.Logger
var once sync.Once

func InitLogger() {
	once.Do(func() {
		if config.IsProduction() {
			logger, _ = zap.NewProduction()
		} else {
			logger, _ = zap.NewDevelopment()
		}
		defer logger.Sync()
	})
}

func Client() *zap.Logger {
	return logger
}
