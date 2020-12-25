package DCEAPI

type SpotAPI interface {
	LimitBuyOrder()
	LimitSellOrder()
	MarketBuyOrder()
	MarketSellOrder()
	FetchTicker()
}



