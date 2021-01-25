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

type Balance struct {
	Currency string
	available float32
	frozen float32
}

type BaseResponse struct {
	Errno int 
	Errmsg string 
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
func sign(apikey, secret, nonce string, data map[string]string) string {
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
	io.WriteString(t,hashString);
	return fmt.Sprintf("%x",t.Sum(nil));
}

// 生成随机字符串
func getRandomStr() string {
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
func generateNonce() string {
	ts := strconv.FormatInt(time.Now().Unix(),10)
	return ts + "_" + getRandomStr()
}

// 生成请求头
func generateHeader(apikey, secret string, params map[string]string) map[string]string {
	nonce := generateNonce()
	return map[string]string{
		"Nonce": nonce,
		"Token": apikey,
		"Signature": sign(apikey, secret, nonce, params),
	}
}

func (aofex *Aofex) GetExchangeName() string {
	return "AOFEX"
}

// 公共接口请求方法
func (aofex *Aofex) RequestPublic(method, path string, params, body map[string]string) (string, error) {
	res, err := DCEAPI.BaseRequest(method, path, params, body, nil)
	return res, err
}

// 私有接口带请求头请求
func (aofex *Aofex) RequestPrivate(method, path string, params, body map[string]string) (string, error) {
	headers := generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, params)
	res, err := DCEAPI.BaseRequest(method, path, params, body, headers)
	return res, err
}

func (aofex *Aofex) FetchBalance() (*DCEAPI.Balance, error) {
	var err error
	balance := &DCEAPI.Balance{}
	res, err := aofex.RequestPrivate("GET", aofex.SpotPath+"/openApi/wallet/list", map[string]string{"show_all":"1"}, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	err = json.Unmarshal([]byte(res), balance)
	return balance, err
}

func (aofex *Aofex) FetchMarkets() (string, error) {
	var err error
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/symbols", nil, nil)
	return res, err
}

func (aofex *Aofex) FetchTrades(symbol string) (string, error) {
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/trade", map[string]string{"symbol":symbol}, nil)
	return res, err
}

func (aofex *Aofex) FetchDepth(symbol string) (string, error) {
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/depth", map[string]string{"symbol":symbol}, nil)
	return res, err
}

func (aofex *Aofex) FetchOHLCV(symbol, period string, size int) (string, error) {
	params := map[string]string{"symbol":symbol, "period":period, "size":strconv.Itoa(size)}
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/kline", params, nil)
	return res, err
}

func (aofex *Aofex) FetchPercision(symbols ...string) (string, error) {
	params := map[string]string{}
	if len(symbols) != 0 {
		params["symbol"] = symbols[0]
	}
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/precision", params, nil)
	return res, err 
}

func (aofex *Aofex) FetchTicker24H(symbols ...string) (string, error) {
	params := map[string]string{}
	if len(symbols) != 0 {
		params["symbol"] = symbols[0]
	}
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/24kline", params, nil)
	return res, err 
}

func (aofex *Aofex) LimitSellOrder(symbol string, amount, price float64) (string, error) {
	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/add", nil, map[string]string{
		"symbol": symbol,
		"type": "sell-limit",
		"amount": strconv.FormatFloat(float64(amount), 'E', -1, 64),
		"price": strconv.FormatFloat(float64(price), 'E', -1, 64),
	})
	return res, err
}

func (aofex *Aofex) LimitBuyOrder(symbol string, amount, price float64) (string, error) {
	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/add", nil, map[string]string{
		"symbol": symbol,
		"type": "buy-limit",
		"amount": strconv.FormatFloat(float64(amount), 'E', -1, 64),
		"price": strconv.FormatFloat(float64(price), 'E', -1, 64),
	})
	return res, err
}

func (aofex *Aofex) MarketBuy(symbol string, amount float64) (string, error) {
	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/add", nil, map[string]string{
		"symbol": symbol,
		"type": "buy-market",
		"amount": strconv.FormatFloat(float64(amount), 'E', -1, 64),
	})
	return res, err
}

func (aofex *Aofex) MarketSell(symbol string, amount float64) (string, error) {
	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/add", nil, map[string]string{
		"symbol": symbol,
		"type": "sell-market",
		"amount": strconv.FormatFloat(float64(amount), 'E', -1, 64),
	})
	return res, err
}

func (aofex *Aofex) CancelOrderBySymbol(symbol string) (string, error) {
	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/cancel", nil, map[string]string{
		"symbol": symbol,
	})
	return res, err
}

func (aofex *Aofex) CancelOrderByID(orderIDs ...string) (string, error) {
	var orderIDsStr string
	for _, v := range orderIDs {
		orderIDsStr += v
	}
	res, err := aofex.RequestPrivate("POST", aofex.SpotPath+"/openApi/entrust/cancel", nil, map[string]string{
		"symbol": orderIDsStr,
	})
	return res, err
}

func (aofex *Aofex) FetchCurrentOrders() (string, error) {
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/entrust/currentList", nil, nil)
	return res, err
}
