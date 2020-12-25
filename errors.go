package DCEAPI

import (
	"fmt"
)

type ExchangeError struct {
	err_code int
	err_msg string
}

func (exErr ExchangeError) Error() ExchangeError {
	fmt.Println(exErr.err_msg)
	return exErr
}