package DCEAPI

type SpotAPI interface {
	GetExchangeName() string
	FetchBalance() ([]Balance, error)
	FetchMarkets() ([]Market, error)
	FetchTrades(symbol string) (Trade, error)
	FetchDepth(symbol string) (OrderBook, error)
	FetchOHLCV(symbol, period string, size int) (Kline, error)
	FetchPercision(symbols ...string) (Precision, error)
	FetchKline24H(symbols ...string) ([]Kline24H, error)
	LimitBuyOrder(symbol, amount, price string) (Order, error)
	LimitSellOrder(symbol, amount, price string) (Order, error)
	MarketSellOrder(symbol, amount, price string) (Order, error)
	MarketBuyOrder(symbol, amount, price string) (Order, error)
	CancelOrderByIDs(orderIDs ...string) ([]Order, []Order, error)
	CancelOrderBySymbol(symbol string) ([]Order, []Order, error)
	FetchOpenOrders() ([]Order, error)
	FetchClosedOrder() ([]Order, error)
	FetchOrder(OrderID string) (Order, error)
}

type SwapAPI interface {
	GetExchangeName() string
}



