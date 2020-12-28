package DCEAPI

type SpotAPI interface {
	LimitBuyOrder()
	LimitSellOrder()
	MarketBuyOrder()
	MarketSellOrder()
	FetchTicker()
	FetchBalance()
}

type SwapAPI interface {
	SwapLimitBuy()
	SwapLimitSell()
}


