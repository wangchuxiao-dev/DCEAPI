package DCEAPI

import (
	_"encoding/json"
	"github.com/shopspring/decimal"
)

// 交易对结构
type Symbol struct {
	Symbol string
	Base string
	Quote string
}

// 市场结构体
type Market struct {
	Symbol string
	BaseCurrency string
	QuoteCurrency string
	PricePrecision int
	AmountPrecision int
	LimitMinOrderAmount float64 
	LimitMaxOrderAmount float64 
	SellMarketMaxOrderAmount float64
	SellMarketMinOrderAmount float64 
	BuyMarketMaxValue float64 
	MinOrderValue float64
}

// 余额结构体
type Balance struct {
	Currency string
	Free decimal.Decimal
	Frozen decimal.Decimal
}

// 市场成交结构体
type Trade struct {
	Symbol string
	Amount decimal.Decimal
	Price decimal.Decimal
	Side string
	Ts int
}

// orderbook
type OrderBook struct {
	Symbol string
	Ts int
	Bids, Asks [][2]decimal.Decimal
}

// k线结构体
type Kline struct {
	Amount decimal.Decimal
	Count decimal.Decimal
	Open, Close, Low, High, Vol decimal.Decimal	
}


// 订单结构体
type Order struct {
	OrderID string //订单号
	Symbol string //交易对
	CreateTime int //创建时间
	ClosedTime int //成交时间
	Type string //类型 market/limit
	Side string //方向 buy/sell
	Price decimal.Decimal // 下单价格 
	Amount decimal.Decimal // 下单数量
	DealPrice decimal.Decimal // 实际成交价格
	FilledAmountQuote decimal.Decimal //已经成交数量报价
	FilledAmountBase decimal.Decimal // 已经成交数量基准v
	Fee decimal.Decimal //手续费
	Status string //订单状态
}