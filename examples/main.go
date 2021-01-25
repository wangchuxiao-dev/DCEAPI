 package main

import (
	"github.com/PythonohtyP1900/DCEAPI/exchanges"
	"fmt"
)

func main() {
	token := "4bec6394e490aca7acaae197379824d3"
	secret := "k51r7mii94jlebhk4ahq"
	
	aofex := exchanges.NewAofex(secret, token)
	balance, _ := aofex.FetchBalance()
	fmt.Println(balance.Result)
	// fmt.Println(balance)

	// trades, _ := aofex.FetchTrades("BTC-USDT")
	// fmt.Println(trades)

	// depth, _ := aofex.FetchDepth("BTC-USDT")
	// fmt.Println(depth)

	// kline, _ := aofex.FetchOHLCV("BTC-USDT", "1min", 100)
	// fmt.Println(kline)

	// pre, _ := aofex.FetchPercision("BTC-USDT")
	// fmt.Println(pre)
}