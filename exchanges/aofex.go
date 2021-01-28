package exchanges

import (
	"crypto/sha1"
	"math/rand" 
	"fmt"
	"sort"
	"strconv"
	"io"
	"encoding/json"
	"time"

	"github.com/PythonohtyP1900/DCEAPI"
)

const (
	PATH string = "https://aofex.com/"
	CNPATH string = "https://aofex.co/"
	SPOTPATH string = "https://openapi.aofex.co"
	SWAPPATH string = "https://openapi-contract.aofex.co"
	DEBUGSPOTPATH string = "https://openapi.aofex.co"
	DEBUGSWAPPATH string = "https://openapi-contract.aofex.co"
)

type Aofex struct {
	Path string
	SpotPath string
	SwapPath string
	Exchange *DCEAPI.Exchange
}		

type BaseResponse struct {
	Errno int 
	ErrMsg string 
}

func NewAofex(secret, apikey string) *Aofex {
	aofex := &Aofex{
		Path: PATH,
		SpotPath: SPOTPATH,
		SwapPath: SWAPPATH,
		Exchange: &DCEAPI.Exchange{
			Name: "AOFEX",
			Secret: secret,
			Apikey: apikey,
		},
	}
	return aofex
}

func DebugNewAofex(secret, apikey string) *Aofex {
	aofex := &Aofex{
		Path: PATH,
		SpotPath: DEBUGSPOTPATH,
		SwapPath: DEBUGSWAPPATH,
		Exchange: &DCEAPI.Exchange{
			Name: "AOFEX",
			Secret: secret,
			Apikey: apikey,
		},
	}
	return aofex
}

// 签名
func (aofex *Aofex) sign(apikey, secret, nonce string, data map[string]string) string {
	tmp := sort.StringSlice{apikey, secret, nonce}
	for k, v := range data {
		tmp = append(tmp, k+"="+v)
	}
	sort.Sort(tmp)
	var hashString string
	for _, v := range tmp {
		hashString += v
	}
	t := sha1.New();
	io.WriteString(t, hashString);
	return fmt.Sprintf("%x",t.Sum(nil));
}

// 生成随机字符串
func (aofex *Aofex) getRandomStr() string {
	rand.Seed(time.Now().UnixNano())
	str := "abcdefghijklmnopqrstuvwxyz"
	var randStr string
	for i:=0;i<4;i++ {
		randInt := rand.Intn(len(str)-1)
		randStr += string(str[randInt])
	}
	return randStr + strconv.Itoa(rand.Intn(9))
}

// 时间戳+随机字符串
func (aofex *Aofex) generateNonce() string {
	ts := strconv.FormatInt(time.Now().Unix(),10)
	return ts + "_" + aofex.getRandomStr()
}

// 生成请求头
func (aofex *Aofex) generateHeader(apikey, secret string, params map[string]string) map[string]string {
	nonce := aofex.generateNonce()
	return map[string]string{
		"Nonce": nonce,
		"Token": apikey,
		"Signature": aofex.sign(apikey, secret, nonce, params),
	}
}

func (aofex *Aofex) GetExchangeName() string {
	return "AOFEX"
}

func (aofex *Aofex) handleError(baseResponse *BaseResponse) error {
	var err error
	switch baseResponse.Errno {
		case 0:
			err = nil
		case 20501, 20502:
			err = &DCEAPI.SymbolError{ErrCode:20501, ErrMsg:baseResponse.ErrMsg}
		case 20506:
			err = &DCEAPI.InvalidSignatureError{ErrCode:20506, ErrMsg:baseResponse.ErrMsg}
		case 20510, 20511, 20512, 20513, 20514:
			err = &DCEAPI.OrderLimitError{ErrCode:baseResponse.Errno, ErrMsg:baseResponse.ErrMsg}
		case 20515:
			err = &DCEAPI.OrderNotFound{ErrCode:20515, ErrMsg:baseResponse.ErrMsg}
		case 20516:
			err = &DCEAPI.OrderStateError{ErrCode:20516, ErrMsg:baseResponse.ErrMsg}
		default:
			err = &DCEAPI.ExchangeError{ErrCode:baseResponse.Errno, ErrMsg:baseResponse.ErrMsg}
	}
	return err
}

func (aofex *Aofex) request(method, path string, params map[string]string, body interface{}, headers map[string]string, model interface{}) error {
	path = DCEAPI.BuildRequestUrl(path, params)
	body_str, err := DCEAPI.BuildRequestBody(body)
	if err != nil {
		return err
	}
	res, err := DCEAPI.HttpRequest(method, path, body_str, headers)
	if err != nil {
		return err
	}
	baseResponse := &BaseResponse{}
	json.Unmarshal(res, baseResponse)
	//如果返回的json errno不等于0就代表请求错误
	exchangeErr := aofex.handleError(baseResponse)
	if exchangeErr != nil {
		return exchangeErr
	}
	return json.Unmarshal(res, model)
}

func (aofex *Aofex) FetchBalance() (*DCEAPI.Balance, error) {
	balance := &DCEAPI.Balance{}
	params := map[string]string{"show_all":"1"}
	headers := aofex.generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, params)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/wallet/list", params, nil, headers, balance)
	return balance, err
}

func (aofex *Aofex) FetchMarkets() (*DCEAPI.Market, error) {
	market := &DCEAPI.Market{}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/symbols", nil, nil, nil, market)
	return market, err
}

func (aofex *Aofex) FetchTrades(symbol string) (*DCEAPI.Trade, error) {
	trade := &DCEAPI.Trade{}
	params := map[string]string{"symbol":symbol}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/trade", params, nil, nil, trade)
	return trade, err
}

func (aofex *Aofex) FetchDepth(symbol string) (*DCEAPI.OrderBook, error) {
	orderbook := &DCEAPI.OrderBook{}
	params := map[string]string{"symbol":symbol}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/depth", params, nil, nil, orderbook)
	return orderbook, err
}

func (aofex *Aofex) FetchOHLCV(symbol, period string, size int) (*DCEAPI.Kline, error) {
	kline := &DCEAPI.Kline{}
	params := map[string]string{"symbol":symbol, "period":period, "size":strconv.Itoa(size)}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/kline", params, nil, nil, kline)
	return kline, err
}

func (aofex *Aofex) FetchPercision(symbols ...string) (*DCEAPI.Precision, error) {
	precision := &DCEAPI.Precision{}
	params := map[string]string{}
	if len(symbols) != 0 {
		params["symbol"] = symbols[0]
	}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/precision", params, nil, nil, precision)
	return precision, err 
}

func (aofex *Aofex) FetchKline24H(symbols ...string) (*DCEAPI.Kline24H, error) {
	kline := &DCEAPI.Kline24H{}
	params := map[string]string{}
	if len(symbols) != 0 {
		params["symbol"] = symbols[0]
	}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/24kline", params, nil, nil, kline)
	return kline, err 
}

func (aofex *Aofex) Order(side, symbol, amount, price string) (*DCEAPI.Order, error) {
	order := &DCEAPI.Order{}
	body := map[string]string{
		"symbol": symbol,
		"type": side,
		"amount": amount,
		"price": price,
	 }
	headers := aofex.generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, body)
	err := aofex.request("POST", aofex.SpotPath+"/openApi/entrust/add", nil, body, headers, order)
	return order, err
}

// func (aofex *Aofex) LimitSellOrder(symbol string, amount, price float64) (string, error) {
// 	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/add", nil, map[string]string{
// 		"symbol": symbol,
// 		"type": "sell-limit",
// 		"amount": strconv.FormatFloat(float64(amount), 'E', -1, 64),
// 		"price": strconv.FormatFloat(float64(price), 'E', -1, 64),
// 	})
// 	return res, err
// }

// func (aofex *Aofex) LimitBuyOrder(symbol string, amount, price float64) (string, error) {
// 	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/add", nil, map[string]string{
// 		"symbol": symbol,
// 		"type": "buy-limit",
// 		"amount": strconv.FormatFloat(float64(amount), 'E', -1, 64),
// 		"price": strconv.FormatFloat(float64(price), 'E', -1, 64),
// 	})
// 	return res, err
// }

// func (aofex *Aofex) MarketBuy(symbol string, amount float64) (string, error) {
// 	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/add", nil, map[string]string{
// 		"symbol": symbol,
// 		"type": "buy-market",
// 		"amount": strconv.FormatFloat(float64(amount), 'E', -1, 64),
// 	})
// 	return res, err
// }

// func (aofex *Aofex) MarketSell(symbol string, amount float64) (string, error) {
// 	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/add", nil, map[string]string{
// 		"symbol": symbol,
// 		"type": "sell-market",
// 		"amount": strconv.FormatFloat(float64(amount), 'E', -1, 64),
// 	})
// 	return res, err
// }

// func (aofex *Aofex) CancelOrderBySymbol(symbol string) (string, error) {
// 	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/cancel", nil, map[string]string{
// 		"symbol": symbol,
// 	})
// 	return res, err
// }

// func (aofex *Aofex) CancelOrderByID(orderIDs ...string) (string, error) {
// 	var orderIDsStr string
// 	for _, v := range orderIDs {
// 		orderIDsStr += v
// 	}
// 	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/cancel", nil, map[string]string{
// 		"symbol": orderIDsStr,
// 	})
// 	return res, err
// }

// func (aofex *Aofex) FetchCurrentOrders() (string, error) {
// 	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/entrust/currentList", nil, nil)
// 	return res, err
// }
