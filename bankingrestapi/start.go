package main

import (
	"example.com/banking/controllers"
	"github.com/gin-gonic/gin"
)

const (
	AuthLogin    = "admin"
	AuthPassword = "password"
)

func main() {

	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		AuthLogin: AuthPassword,
	}))

	authorized.GET("/api/v1/wallets/:id/balance", controllers.GetBalance)
	authorized.POST("/api/v1/wallets/:id/credit", controllers.CreditWallet)
	authorized.POST("/api/v1/wallets/:id/debit", controllers.DebitWallet)

	r.Run()
}
