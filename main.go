package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Tonkpils/lendingclub"
	"github.com/shopspring/decimal"
)

func main() {
	apiKey := os.Getenv("LC_KEY")
	accountID, err := strconv.Atoi(os.Getenv("LC_ACCOUNT_ID"))
	if err != nil {
		log.Fatal("Error reading LendingClub Account ID.")
	}

	c := lendingclub.NewClient(apiKey, nil)
	ar := c.Accounts(accountID)
	ac, err := ar.AvailableCash()
	if err != nil {
		log.Fatal(err)
	}

	zero := decimal.New(0, 0)

	if ac.AvailableCash.GreaterThan(zero) {
		wd, err := ar.WithdrawFunds(ac.AvailableCash)
		fmt.Printf("Withdrawing %v. Expected Delivery: %s", wd.Amount, wd.EstimatedFundsTransferDate)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("No available cash. Aborting.")
	}
}
