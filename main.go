package main

import (
  "context"
  "fmt"
  "github.com/adshao/go-binance/v2"
  "github.com/guozhe001/ex-api-go/config"
  "github.com/guozhe001/ex-api-go/constant"
  "time"
)

func main() {
  apiKey, secret := GetApiKeyAndSecret(constant.ExchangeBinance)
  fmt.Printf("apiKey=%s, secret=%s\n", apiKey, secret)
  client := binance.NewClient(apiKey, secret)
  date := time.Date(2021, 4, 19, 0, 0, 0, 0, time.Local)
  startTime := date.UnixNano() / 1000 / 1000
  trades, err := client.NewAggTradesService().
    Symbol("BNBBUSD").
    StartTime(startTime).
    EndTime(startTime + 30*60*1000).
    Do(context.Background())
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, t := range trades {
    fmt.Println(t)
  }
}

// GetApiKeyAndSecret 获取指定交易所的apikey和secret
// ex 交易所
func GetApiKeyAndSecret(ex string) (string, string) {
  apiKey, secret := config.Get(ex, constant.ApiKey), config.Get(ex, constant.Secret)
  return apiKey, secret
}
