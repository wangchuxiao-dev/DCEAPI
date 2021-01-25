package DCEAPI

type WalletApi interface {
	FetchBalance() (string, error)
}

type SpotAPI interface {
	GetExchangeName() string
	FetchMarkets() (string, error)
	LimitBuyOrder(symbol string, amount, price float64) (string, error)
	LimitSellOrder(symbol string, amount, price float64) (string, error)
	MarketBuyOrder(symbol string, amount float64) (string, error)
	MarketSellOrder(symbol string, amount float64) (string, error)
	FetchTicker() (string, error)
	FetchTicker24H() (string, error)
}

type SwapAPI interface {
	SwapLimitBuy()
	SwapLimitSell()
	SwapMarketBuy()
	SwapMarketSell()
}


