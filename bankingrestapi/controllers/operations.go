package controllers

import (
	"net/http"
	"strconv"

	"example.com/banking/datalib"
	"example.com/banking/helpers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

func GetBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting walletId: " + c.Param("id")})
		return
	}

	json, errDb := datalib.GetWalletById(id)
	if errDb != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Redis error occured trying to get the data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": json})
}

func CreditWallet(c *gin.Context) {
	amountStr := c.PostForm("amount")
	amountAsFloat, errAmount := strconv.ParseFloat(amountStr, 64)
	if errAmount != nil {
		Logger.Errorf("Error converting amount %v", errAmount)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting amount"})
		return
	}

	if amountAsFloat < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error: negative amount"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting walletId"})
		return
	}

	json, errDb := datalib.GetWalletById(id)
	if errDb != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Redis error occured trying to get the data"})
		return
	}

	model := helpers.ConvertJsonToWallet(json)
	model.Balance += amountAsFloat

	errDb = datalib.SetWallet(model)

	if errDb != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Redis error updating the data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "OK"})
}

func DebitWallet(c *gin.Context) {
	amountStr := c.PostForm("amount")
	amountAsFloat, errAmount := strconv.ParseFloat(amountStr, 64)
	if errAmount != nil {
		Logger.Errorf("Error converting amount %v", errAmount)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting amount"})
		return
	}

	if amountAsFloat < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error: negative amount"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting walletId"})
		return
	}

	json, errDb := datalib.GetWalletById(id)
	if errDb != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Redis error occured trying to get the data"})
		return
	}

	model := helpers.ConvertJsonToWallet(json)
	if model.Balance < amountAsFloat {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Balance is less than transaction's request"})
		return
	}

	model.Balance -= amountAsFloat
	errDb = datalib.SetWallet(model)

	if errDb != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Redis error updating the data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "OK"})
}
