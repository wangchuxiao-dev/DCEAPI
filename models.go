package DCEAPI

import (
	_"encoding/json"
)

// 市场结构体
type Market struct {
	Symbol string
	Base_currency string
	Quote_currency string
	Min_size float64
	Max_size float64
	Max_price float64
	Min_pirce float64
	Maker_fee float64
	Taker_fee float64
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
	Order_id int
	Order_sn string
	Symbol string
	Ctime string 
	Type int
	Side string 
	Price string 
	Number string 
	Total_price string
	Deal_number string 
	Deal_price string
	Status int
}