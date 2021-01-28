package DCEAPI

import (
	"fmt"
)

// http状态码!=200时候抛出
type HttpError struct {
	HttpCode int
	HttpMsg string
}

// 认证错误
type InvalidSignatureError struct {
	ErrCode int
	ErrMsg string
}

// 交易对错误
type SymbolError struct {
	ErrCode int
	ErrMsg string
}

// 余额不足
type BalanceError struct {
	ErrCode int
	ErrMsg string
}

// 未知订单
type OrderNotFound struct {
	ErrCode int
	ErrMsg string
}

// 交易所错误
type ExchangeError struct {
	ErrCode int
	ErrMsg string
}

// 下单精度，最大最小量错误
type OrderLimitError struct {
	ErrCode int
	ErrMsg string
}

type OrderStateError struct {
	ErrCode int
	ErrMsg string
}

func (exErr InvalidSignatureError) Error() string {
	return fmt.Sprintf("invalid signature, error_code:%d, error_message:%s", exErr.ErrCode, exErr.ErrMsg)
}

func (exErr ExchangeError) Error() string {
	return fmt.Sprintf("Exchange error, error_code:%d, error_message:%s", exErr.ErrCode, exErr.ErrMsg)
}

func (exErr OrderNotFound) Error() string {
	return fmt.Sprintf("invalid order id, error_code:%d, error_message:%s", exErr.ErrCode, exErr.ErrMsg)
}

func (exErr BalanceError) Error() string {
	return fmt.Sprintf("invalid order id, error_code:%d, error_message:%s", exErr.ErrCode, exErr.ErrMsg)
}

func (exErr SymbolError) Error() string {
	return fmt.Sprintf("invalid symbol, error_code:%d, error_message:%s", exErr.ErrCode, exErr.ErrMsg)
}

func (exErr OrderLimitError) Error() string {
	return fmt.Sprintf("order limit error, error_code:%d, error_message:%s", exErr.ErrCode, exErr.ErrMsg)
}

func (exErr OrderStateError) Error() string {
	return fmt.Sprintf("order state error, error_code:%d, error_message:%s", exErr.ErrCode, exErr.ErrMsg)
}