package DCEAPI

type Ticker struct {
	Symbol string
	Last float64      
	Buy  float64      
	Sell float64     
	High float64      
	Low  float64      
	Vol  float64 
	Date uint64      
}

type Balance struct {
	Currency string
	available float64
	frozen float64
}

type Symbol struct {
	Symbol string
	Base string
	Quoto string
}

type Order struct {
	ID int
	Side string
	Amount float64
	Price float64
	TotalPrice float64
	Timestamp int
}

