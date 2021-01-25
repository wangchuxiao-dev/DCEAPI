package DCEAPI

import (
	"fmt"
)

// 余额不足
type BalanceError struct {
	err_code int
	err_msg string
}

// 未知订单
type OrderNotFound struct {
	err_code int
	err_msg string
}

// 交易所错误
type ExchangeError struct {
	err_code int
	err_msg string
}

func (exErr ExchangeError) Error() string {
	return fmt.Sprintf("error")
}