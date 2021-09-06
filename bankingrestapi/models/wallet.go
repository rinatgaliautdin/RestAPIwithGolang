package models

type Wallet struct {
	Id      int     `json:"id" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}
