package DCEAPI

import (
	_ "github.com/shopspring/decimal"
	_"encoding/json"
)

// k线结构体
type Ticker struct {
	Result []struct{
		Symbol string
		Last float64      
		Buy  float64      
		Sell float64     
		High float64      
		Low  float64      
		Vol  float64 
		Date uint64    
	}  
}

// 余额结构体
type Balance struct {
	Result []struct{
		Currency string
		Available float64
		Frozen float64
	}
}

// 交易对结构体
type Symbol struct {
	Symbol string
	Base string
	Quoto string
}

// 订单结构体
type Order struct {
	ID int
	Side string
	Amount float64
	Price float64
	TotalPrice float64
	Timestamp int
}

