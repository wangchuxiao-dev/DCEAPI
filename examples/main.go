 package main

import (
	"github.com/PythonohtyP1900/DCEAPI/exchanges"
	"github.com/PythonohtyP1900/DCEAPI"
	"fmt"
)

func main() {
	token := ""
	secret := ""

	aofex := exchanges.NewAofex(secret, token)
	var spotApi DCEAPI.SpotAPI
	spotApi = aofex

	balance, err := spotApi.FetchBalance()
	fmt.Println(balance, err)

	markets, err := spotApi.FetchMarkets()
	fmt.Println(markets, err)

	trades, err := spotApi.FetchTrades("BTC-USDT")
	fmt.Println(trades, err)

	depth, err := spotApi.FetchDepth("BTC-USDT")
	fmt.Println(depth, err)

	kline, err := spotApi.FetchOHLCV("BTC-USDT", "5min", 100)
	fmt.Println(kline, err)

	precision, err := spotApi.FetchPercision("BTC-USDT")
	fmt.Println(precision, err)

	kline24, err := spotApi.FetchKline24H("BTC-USDT")
	fmt.Println(kline24, err)

	order, err := spotApi.LimitBuyOrder("EOS-AQ","2", "0.88")
	fmt.Println(order, err)

	result, err := spotApi.FetchOrder(order.Order_sn)
	fmt.Println(result, err)

	sucess, failed, err := spotApi.CancelOrderByIDs(order.Order_sn)
	fmt.Println(sucess, failed, err)

	orders, err := spotApi.FetchClosedOrder()
	fmt.Println(orders, err)
}
