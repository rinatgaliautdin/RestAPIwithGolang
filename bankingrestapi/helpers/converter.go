package helpers

import (
	"encoding/json"

	"example.com/banking/models"
)

func ConvertJsonToWallet(jsonStr string) *models.Wallet {
	model := models.Wallet{}
	json.Unmarshal([]byte(jsonStr), &model)

	return &model
}
