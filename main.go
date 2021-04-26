package main

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/guozhe001/ex-api-go/config"
	"github.com/guozhe001/ex-api-go/constant"
	"strconv"
)

func main() {
	apiKey, secret := GetApiKeyAndSecret(constant.ExchangeBinance)
	fmt.Printf("apiKey=%s, secret=%s\n", apiKey, secret)
	client := binance.NewClient(apiKey, secret)
	balances, err := getAccount(client)
	if err != nil {
		panic(err)
	}
	for _, b := range balances {
		fmt.Print(b)
	}
}

// GetApiKeyAndSecret 获取指定交易所的apikey和secret
// ex 交易所
func GetApiKeyAndSecret(ex string) (string, string) {
	apiKey, secret := config.Get(ex, constant.ApiKey), config.Get(ex, constant.Secret)
	return apiKey, secret
}

func getAccount(client *binance.Client) ([]binance.Balance, error) {
	var balance []binance.Balance
	response, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return balance, err
	}
	fmt.Printf("response.CanDeposit=%t, response.CanTrade=%t, response.CanWithdraw=%t, response.BuyerCommission=%d,"+
		" response.MakerCommission=%d, response.SellerCommission=%d, response.TakerCommission=%d\n", response.CanDeposit,
		response.CanTrade, response.CanWithdraw, response.BuyerCommission,
		response.MakerCommission, response.SellerCommission, response.TakerCommission)
	for _, item := range response.Balances {
		free, err := strconv.ParseFloat(item.Free, 64)
		if err != nil {
			return balance, err
		}
		locked, err := strconv.ParseFloat(item.Locked, 64)
		if err != nil {
			return balance, err
		}
		if free+locked > 0 {
			fmt.Printf("%#v\n", item)
			balance = append(balance, item)
		}
	}
	return balance, nil
}
