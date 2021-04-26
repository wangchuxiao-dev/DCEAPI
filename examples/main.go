 package main

import (
	"github.com/PythonohtyP1900/DCEAPI/exchanges"
	"fmt"
)

func main() {
	apiKeyHB := "9a159d7e-vfd5ghr532-b48227db-511af"
	secretHB := "b59ac6d6-f6d0b057-a088759f-87368"

	huobi := exchanges.NewHuobi(secretHB, apiKeyHB)
	order, err := huobi.LimitBuyOrder("btcusdt", "0.02", "52000")
	fmt.Println(order, err)

	order, err = huobi.LimitSellOrder("btcusdt", "0.02", "52000")
	fmt.Println(order, err)

	order, err = huobi.MarketBuyOrder("btcusdt", "10000")
	fmt.Println(order, err)

	order, err = huobi.MarketSellOrder("btcusdt", "0.9")
	fmt.Println(order, err)
	// apiKeyAOFEX := "348573c7a7b3d0f002fbea148cdf8571"
	// secretAOFEX := "f6vwxrqoz2ptkrgb9t0m"
	// aofex := exchanges.NewAofex(secretAOFEX, apiKeyAOFEX)
	// order, err := aofex.LimitBuyOrder("BTC-USDT", "0.01", "51000")
	// fmt.Println(order, err)
}
