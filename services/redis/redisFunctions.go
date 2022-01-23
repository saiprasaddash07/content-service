package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/saiprasaddash07/content-service.git/config"
	"github.com/saiprasaddash07/content-service.git/constants"
	"github.com/saiprasaddash07/content-service.git/services/logger"
	"go.uber.org/zap"
)

func Get(key string) (string, error) {
	if !config.Get().CacheEnabled {
		return "", errors.New(constants.INFO_CACHE_DISABLED)
	}
	val, err := GetClient().Get(key).Result()
	if err == redis.Nil {
		logger.Client().Info("NOT FOUND----->", zap.Error(err))
		return "", err
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func Set(key string, value string, duration time.Duration) (string, error) {
	if !config.Get().CacheEnabled {
		return "", errors.New(constants.INFO_CACHE_DISABLED)
	}
	val, err := GetClient().Set(key, value, duration).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func Delete(key string) (bool, error) {
	if !config.Get().CacheEnabled {
		return true, errors.New(constants.INFO_CACHE_DISABLED)
	}
	var _, err = GetClient().Del(key).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}

func Increment(key string) (int64, error) {
	if !config.Get().CacheEnabled {
		return 0, errors.New(constants.INFO_CACHE_DISABLED)
	}
	val, err := GetClient().Del(key).Result()
	if err == redis.Nil {
		logger.Client().Info("NOT FOUND----->", zap.Error(err))
	} else if err != nil {
		logger.Client().Info("Err in redis-->", zap.Error(err))
	}
	return val, nil
}
