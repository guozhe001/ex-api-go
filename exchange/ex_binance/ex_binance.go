package ex_binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/guozhe001/ex-api-go/config"
	"github.com/guozhe001/ex-api-go/constant"
	"log"
	"strconv"
)

// GetApiKeyAndSecret 获取指定交易所的apikey和secret
// ex 交易所
func GetApiKeyAndSecret(ex string) (string, string) {
	apiKey, secret := config.Get(ex, constant.ApiKey), config.Get(ex, constant.Secret)
	return apiKey, secret
}

// 获取币安的client
func NewClient() *binance.Client {
	apiKey, secret := GetApiKeyAndSecret(constant.ExchangeBinance)
	fmt.Printf("apiKey=%s, secret=%s\n", apiKey, secret)
	return binance.NewClient(apiKey, secret)
}

func GetAccount(client *binance.Client) ([]binance.Balance, error) {
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

// GET /api/v3/ticker/price
func GetTickerPrice(client *binance.Client, symbol string) {
	service := client.NewListPricesService()
	if symbol != "" {
		service.Symbol(symbol)
	}
	res, err := service.Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range res {
		log.Print(p)
	}
}
