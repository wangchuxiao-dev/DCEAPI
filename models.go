package DCEAPI

type Symbol struct {
	Symbol string
	Base string
	Quoto string
}

type Ticker struct {
	Open float64
	Close float64
	High float64
	Low float64
	Volume float64
	Timestamp int
}

type Order struct {
	ID int
	Side string
	Amount float64
	Price float64
	TotalPrice float64
	Timestamp int
}