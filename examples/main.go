 package main

import (
	"github.com/PythonohtyP1900/DCEAPI/exchanges"
	_ "github.com/shopspring/decimal"

	"fmt"
)

func main() {
	exchangeName := "aofex"
	var apiKey string
	var secret string
	switch exchangeName {
	case "aofex":
		apiKey = "eced48e1-08f6d127-ht4tgq1e4t-6c58d"
		secret = "547aa112-8fdee198-e7753233-707d8"
	case "huobipro":
		apiKey = "eced48e1-08f6d127-ht4tgq1e4t-6c58d"
		secret = "547aa112-8fdee198-e7753233-707d8"
	}
	api := exchanges.NewExchange(exchangeName, apiKey, secret)

	// balance, _ := api.FetchBalance()
	// var usdtBalance decimal.Decimal
	// for _, v := range balance {
	// 	if v.Currency == "trx" {
	// 		usdtBalance = v.Free
	// 	}
	// }
	// fmt.Println(usdtBalance)

	// trade, err := api.FetchTrades("BTC/USDT", "1")
	// fmt.Println(trade, err)

	// orderBook, err := api.FetchOrderBook("BTC/USDT", map[string]string{"type":"step0"})
	// fmt.Println(orderBook, err)

	// kline, err := api.FetchOHLCV("BTC/USDT", "1min", "200")
	// fmt.Println(kline, err)

	// kline24H, err := api.FetchOHLCV24H("BTC/USDT")
	// fmt.Println(kline24H, err)

	markets, err := api.FetchMarkets()
	fmt.Println(markets, err)

	// order, err := huobi.LimitBuyOrder("TRX/USDT", "50", "0.144700")
	// fmt.Println(order, err)
	
	// order, err = huobi.MarketBuyOrder("btcusdt", "10000")
	// fmt.Println(order, err)

	// order, err := huobi.MarketBuyOrder("trxusdt", "5")
	// fmt.Println(order, err)

	// order, err := huobi.MarketBuyOrder("trxusdt", "5")
	// fmt.Println(order, err)

	// order, err := huobi.LimitBuyOrder("TRX/USDT", "50", "0.1247")
	// fmt.Println(order, err)

	// err = huobi.CancelOrder("273169416626268")
	// fmt.Println(err)
	// order, err := huobi.FetchOrder("273169467603314")
	// fmt.Println(order, err)
	// apiKeyAOFEX := "348573c7a7b3d0f002fbea148cdf8571"
	// secretAOFEX := "f6vwxrqoz2ptkrgb9t0m"
	// aofex := exchanges.NewAofex(secretAOFEX, apiKeyAOFEX)
	// balance, _ = aofex.FetchBalance()
	// fmt.Println(balance)
	// order, err = aofex.LimitSellOrder("BTC/USDT", "0.01", "591000")
	// fmt.Println(order, err)
}
