package tests

import (
	"testing"

	"example.com/banking/datalib"

	"example.com/banking/models"
)

func PopulateRedis() error {

	wallet1 := models.Wallet{
		Id:      1,
		Balance: 100,
	}
	wallet2 := models.Wallet{
		Id:      2,
		Balance: 120,
	}
	wallet3 := models.Wallet{
		Id:      3,
		Balance: 130,
	}
	wallet4 := models.Wallet{
		Id:      4,
		Balance: 140,
	}
	wallet5 := models.Wallet{
		Id:      5,
		Balance: 150,
	}

	err := datalib.SetWallet(&wallet1)
	if err != nil {
		return err
	}

	err = datalib.SetWallet(&wallet2)
	if err != nil {
		return err
	}
	err = datalib.SetWallet(&wallet3)
	if err != nil {
		return err
	}
	err = datalib.SetWallet(&wallet4)
	if err != nil {
		return err
	}
	err = datalib.SetWallet(&wallet5)
	if err != nil {
		return err
	}

	return nil
}

func TestDb(t *testing.T) {
	err := PopulateRedis()

	if err != nil {
		t.Errorf("Population of the Db failed")
	}

	currentWallet1, errData := datalib.GetWalletById(1)
	if errData != nil {
		t.Errorf("Error retreiving the wallet")
	}
	t.Logf(currentWallet1)

	if currentWallet1 == "" {
		t.Errorf("Wrong data in the db")
	}
}
