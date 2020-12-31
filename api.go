package DCEAPI

type SpotAPI interface {
	GetExchangeName() string
	FetchBalance() (string, error)
	FetchMarkets() (string, error)
	LimitBuyOrder(symbol string, amount, price float64) (string, error)
	LimitSellOrder(symbol string, amount, price float64) (string, error)
	MarketBuyOrder(symbol string, amount float64) (string, error)
	MarketSellOrder(symbol string, amount float64) (string, error)
	FetchTicker() (string, error) (string, error)
	FetchTicker24H() (string, error) (string, error)
}

type SwapAPI interface {
	SwapLimitBuy()
	SwapLimitSell()
	SwapMarketBuy()
	SwapMarketSell()
}


