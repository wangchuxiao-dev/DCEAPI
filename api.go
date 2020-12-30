package DCEAPI

type SpotAPI interface {
	FetchBalance() (string, error)
	FetchMarkets() (string, error)
	LimitBuyOrder(symbol string, amount, price float64)
	LimitSellOrder(symbol string, amount, price float64)
	MarketBuyOrder(symbol string, amount float64)
	MarketSellOrder(symbol string, amount float64)
	FetchTicker() (string, error)
}

type SwapAPI interface {
	SwapLimitBuy()
	SwapLimitSell()
}


