 package main

import (
	"github.com/PythonohtyP1900/DCEAPI/exchanges"
	"fmt"
)

func main() {
	token := "4bec6394e490aca7acaae197379824d3"
	secret := "k51r7mii94jlebhk4ahq"
	aofex := exchanges.NewAofex(secret, token)

	// balance, err := aofex.FetchBalance()
	// fmt.Println(balance, err)

	// markets, err := aofex.FetchMarkets()
	// fmt.Println(markets, err)

	// trades, err := aofex.FetchTrades("BTC-USDT")
	// fmt.Println(trades, err)

	// depth, err := aofex.FetchDepth("BTC-USDT")
	// fmt.Println(depth, err)

	// kline, err := aofex.FetchOHLCV("BTC-USDT", "5min", 100)
	// fmt.Println(kline, err)

	// precision, err := aofex.FetchPercision("BTC-USDT")
	// fmt.Println(precision, err)

	// kline24, err := aofex.FetchKline24H("BTC-USDT")
	// fmt.Println(kline24, err)

	order, err := aofex.Order("SELL", "BTC-USDT", "33000.89", "2")
	fmt.Println(order, err)
}