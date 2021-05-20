package DCEAPI

// 现货接口
type SpotAPI interface {
	GetExchangeName() string
	// 账户相关接口
	FetchBalance() ([]Balance, error)
	// 行情相关接口
	FetchMarkets() ([]Market, error)
	FetchTrades(symbol, size string) ([]Trade, error)
	FetchTrade(symbol string) ([]Trade, error)
	FetchOrderBook(symbol string, params map[string]string) (OrderBook, error)
	FetchOHLCV(symbol, period, size string) ([]Kline, error)
	FetchOHLCV24H(symbol string) (Kline, error)
	// 订单相关接口
	LimitBuyOrder(symbol, amount, price string) (Order, error)
	LimitSellOrder(symbol, amount, price string) (Order, error)
	MarketSellOrder(symbol, amount string) (Order, error)
	MarketBuyOrder(symbol, amount string) (Order, error)
	FetchOrder(OrderID string) (Order, error)
	CancelOrder(orderID string) (error)
	BatchCancelOrder(orderID ...string) ([]Order, []Order, error)
	CancelOrderBySymbol(symbols ...string) ([]Order, []Order, error)
	FetchOpenOrders(symbols ...string) ([]Order, error)
	FetchClosedOrders(symbols ...string) ([]Order, error)
}

type SwapAPI interface {
	GetExchangeName() string
}



