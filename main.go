package main

import (
	"fmt"
	"github.com/guozhe001/ex-api-go/exchange/ex_binance"
	"log"
)

func main() {
	balances, err := ex_binance.GetAccount(ex_binance.NewClient())
	if err != nil {
		log.Fatal(err)
	}
	for _, b := range balances {
		fmt.Println(b)
	}
}
