package DCEAPI

type ExchangeError struct {
	err_code int
	err_msg string
}

func (exErr ExchangeError) Error() ExchangeError {
	return exErr
}