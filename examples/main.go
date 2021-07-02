package main

import (
	"github.com/PythonohtyP1900/DCEAPI"
	"github.com/PythonohtyP1900/DCEAPI/exchanges"
	"github.com/shopspring/decimal"

	"fmt"
)

func initExchange(exchangeName string) DCEAPI.SpotAPI {
	var apiKey string
	var secret string
	switch exchangeName {
	case "aofex":
		apiKey = ""
		secret = ""
	case "huobipro":
		apiKey = ""
		secret = ""
	}
	api := exchanges.NewExchange(exchangeName, apiKey, secret)
	return api
}

func main() {
	api := initExchange("huobipro")
	balance, err := api.FetchBalance()
	var Free decimal.Decimal
	var Frozen decimal.Decimal
	for _, v := range balance {
		if v.Currency == "USDT" {
			Free = v.Free
			Frozen = v.Frozen
		}
	}
	fmt.Println(Free, Frozen, err)

	trade, err := api.FetchTrades("BTC/USDT", "2")
	fmt.Println(trade, err)

	markets, err := api.FetchMarkets()
	fmt.Println(markets, err)

	orderBook, err := api.FetchOrderBook("BTC/USDT", map[string]string{"type": "step0"})
	fmt.Println(orderBook, err)

	kline, err := api.FetchOHLCV("BTC/USDT", "1min", "200")
	fmt.Println(kline, err)

	kline24H, err := api.FetchOHLCV24H("BTC/USDT")
	fmt.Println(kline24H, err)

	order, err := api.LimitBuyOrder("TRX/USDT", "100", "0.07")
	fmt.Println(order, err)

	orderInfo, err := api.FetchOrder("280985812605956")
	fmt.Println(orderInfo, err)

	err = api.CancelOrder(order.OrderID)
	fmt.Println(err)

	openOrders, err := api.FetchClosedOrders("TRX/USDT")
	fmt.Println(openOrders, err)

	success, failed, err := api.BatchCancelOrder("test_orderid1", "test_orderid2")
	fmt.Println(success, failed, err)

}
