package DCEAPI

type SpotAPI interface {
	GetExchangeName() string
	FetchBalance() ([]Balance, error)
	FetchMarkets() ([]Market, error)
	FetchTrades(symbol, size string) ([]Trade, error)
	FetchTrade(symbol string) ([]Trade, error)
	FetchOrderBook(symbol string, params map[string]string) (OrderBook, error)
	FetchOHLCV(symbol, period, size string) ([]Kline, error)
	FetchOHLCV24H(symbol string) (Kline, error)
	LimitBuyOrder(symbol, amount, price string) (Order, error)
	LimitSellOrder(symbol, amount, price string) (Order, error)
	MarketSellOrder(symbol, amount string) (Order, error)
	MarketBuyOrder(symbol, amount string) (Order, error)
	FetchOrder(OrderID string) (Order, error)
	// CancelOrderByIDs(orderIDs ...string) ([]Order, []Order, error)
	// CancelOrderBySymbol(symbol string) ([]Order, []Order, error)
	// FetchOpenOrders() ([]Order, error)
	// FetchClosedOrder() ([]Order, error)
	// CancelOrder(OrderID string) (Order, error)
}

type SwapAPI interface {
	GetExchangeName() string
}



