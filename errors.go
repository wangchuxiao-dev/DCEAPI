package DCEAPI

import (
	"fmt"
)

type BaseError struct {
	ErrCode int
	ErrMsg string
}

// http状态码!=200时候抛出
type HttpError struct {
	HttpCode int
	HttpMsg string
}

// 认证错误
type InvalidSignatureError BaseError

// 交易对错误
type SymbolError BaseError

// 余额不足
type BalanceError BaseError 

// 未知订单
type OrderNotFound BaseError 

// 交易所错误
type ExchangeError BaseError

// 下单精度，最大最小量错误
type OrderLimitError BaseError

// 状态错误
type OrderStateError BaseError

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
	return fmt.Sprintf("account frozen balance insufficient error, error_code:%d, error_message:%s", exErr.ErrCode, exErr.ErrMsg)
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