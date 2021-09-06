package datalib

import (
	"encoding/json"
	"fmt"

	"example.com/banking/models"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const (
	RedisServer   = "localhost:6379"
	RedisPassword = ""
)

var Logger *logrus.Entry

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisServer,
		Password: RedisPassword,
		DB:       0,
	})

	return client
}

func SetWallet(wallet *models.Wallet) error {
	client := GetRedisClient()
	defer client.Close()

	json, err := json.Marshal(*wallet)
	if err != nil {
		Logger.Errorf("Error: %v", err)
		return err
	}

	var id = fmt.Sprintf("walletId_%d", wallet.Id)
	err = client.Set(id, json, 0).Err()
	if err != nil {
		Logger.Errorf("Error: %v", err)
		return err
	}

	return nil
}

func GetWalletById(id int) (string, error) {
	client := GetRedisClient()
	defer client.Close()

	key := fmt.Sprintf("walletId_%d", id)

	json, err := client.Get(key).Result()
	if err != nil {
		Logger.Errorf("Error: %v", err)
		return "", err
	}

	return json, nil
}
