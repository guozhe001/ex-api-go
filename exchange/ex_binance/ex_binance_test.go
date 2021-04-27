package ex_binance

import (
	"github.com/adshao/go-binance/v2"
	"testing"
)

func TestGetTickerPrice(t *testing.T) {
	type args struct {
		client *binance.Client
		symbol string
	}
	c := NewClient()
	tests := []struct {
		name string
		args args
	}{
		{
			name: "all",
			args: args{client: c, symbol: ""},
		},
		{
			name: "all",
			args: args{client: c, symbol: "BTC-USDT"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetTickerPrice(tt.args.client, tt.args.symbol)
		})
	}
}
