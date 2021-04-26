package DCEAPI

import (
	_"encoding/json"
)

// 市场结构体
type Market struct {
	Symbol string
	BaseCurrency string
	QuoteCurrency string
	MinSize float64
	MaxSize float64
	MaxPrice float64
	MinPrice float64
	MakerFee float64
	TakerFee float64
}

// 余额结构体
type Balance struct {
	Currency string
	Available string
	Frozen string
}

// 市场成交结构体
type Trade struct {	
	Symbol string
	Ts int
	Data []struct{
		Id int
		Amount string
		Price string
		Direction string
		Ts int
	}
}

// orderbook
type OrderBook struct {
	Symbol string
	Ts int
	Bids [][2]string
	Asks [][2]string
}

// 交易对结构体
type Symbol struct {
	Symbol string
	Base string
	Quoto string
}

// k线结构体
type Kline struct {
	Symbol string
	Period string
	Ts int
	Data []struct{
		Id int
		Amount string
		Count float64
		Open, Close, Low, High, Vol string	
	}
}

// Symbol struct{
// 	Amount string
// 	inQuantity string
// 	MaxQuantity string
// 	Price string 
// 	MinPrice string
// 	MaxPrice string 
// }

// 精度结构体
type Precision struct {
	Result map[string]map[string]string
}

// ticker24H结构体
type Kline24H struct {
	Symbol string
	Data struct{
		Id int
		Amount string
		Count float64
		Open, Close, Low, High, Vol string
	}
}

// 订单结构体
type Order struct {
	OrderID string
	Symbol string
	CreateTime int
	ClosedTime int
	Side string 
	Price string 
	Amount string 
	TotalPrice string
	DealAmount string 
	Dealprice string
	Status string
}