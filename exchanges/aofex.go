package exchanges

import (
	"crypto/sha1"
	_ "net/url"
	_ "crypto/rand" 
	"fmt"
	"sort"
	"strconv"
	"io"
	"time"

	"github.com/PythonohtyP1900/DCEAPI"
)

const (
	Path string = "https://aofex.com/"
	CNPath string = "https://aofex.co/"
	SpotPath string = "https://openapi.aofex.co"
	SwapPath string = "https://openapi-contract.aofex.co"
	DebugSpotPath string = ""
	DebugSwapPath string = ""
)

type BaseResponse struct {
	Errno int 
	Errmsg string 
	result interface{}
}

type Aofex struct {
	Path string
	SpotPath string
	SwapPath string
	Exchange *DCEAPI.Exchange
}

func NewAofex(secret, apikey string) *Aofex {
	aofex := &Aofex{
		Path: Path,
		SpotPath: SpotPath,
		SwapPath: SwapPath,
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
		Path: Path,
		SpotPath: DebugSpotPath,
		SwapPath: DebugSwapPath,
		Exchange: &DCEAPI.Exchange{
			Name: "AOFEX",
			Secret: secret,
			Apikey: apikey,
		},
	}
	return aofex
}

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

func generateNonce() string {
	ts := strconv.FormatInt(time.Now().Unix(),10)
	return ts + "_" + "sadx1"
}

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

func (aofex *Aofex) RequestPublic(method, path string, params, body map[string]string) (string, error) {
	res, err := DCEAPI.BaseRequest(method, path, params, map[string]string{}, map[string]string{})
	return res, err
}

func (aofex *Aofex) RequestPrivate(method, path string, params, body map[string]string) (string, error) {
	headers := generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, params)
	res, err := DCEAPI.BaseRequest(method, path, params, nil, headers)
	return res, err
}

func (aofex *Aofex) FetchBalance() (string, error) {
	res, err := aofex.RequestPrivate("GET", aofex.SpotPath+"/openApi/wallet/list", map[string]string{"show_all":"1"}, nil)
	return res, err
}

func (aofex *Aofex) FetchMarkets() (string, error) {
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/symbols", map[string]string{}, map[string]string{})
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

func (aofex *Aofex) FetchOHLCV(symbol string, params map[string]string) (string, error) {
	params["symbol"] = symbol
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/kline", params, nil)
	return res, err
}

func (aofex *Aofex) FetchPercision(symbols ...string) (string, error) {
	var params map[string]string
	if len(symbols) == 0 {
		params = map[string]string{}
	} else {
		params = map[string]string{"symbol":symbols[0]}
	}
	res, err := aofex.RequestPublic("GET", aofex.SpotPath+"/openApi/market/precision", params, nil)
	return res, err 
}

func (aofex *Aofex) FetchTicker24H(symbols ...string) (string, error) {
	var params map[string]string
	if len(symbols) == 0 {
		params = map[string]string{}
	} else {
		params = map[string]string{"symbol":symbols[0]}
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

// func (aofex *Aofex) LimitSellOrder(symbol string, amount, price float64) {

// }

