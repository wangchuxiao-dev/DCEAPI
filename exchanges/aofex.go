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
	"net/url"

	"github.com/PythonohtyP1900/DCEAPI"
)

const (
	PATH string = "https://aofex.com/"
	CNPATH string = "https://aofex.co/"
	SPOTPATH string = "https://openapi.aofex.co"
	SWAPPATH string = "https://openapi-contract.aofex.co"
	DEBUGSPOTPATH string = "http://api.q.xefoa.com/"
	DEBUGSWAPPATH string = "http://openapi.q.xefoa.com"
)

type Aofex struct {
	Path string
	SpotPath string
	SwapPath string
	Exchange *DCEAPI.Exchange
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

type BaseResponse struct {
	Errno int 
	ErrMsg string 
}

type aofexResponse interface {
	hasError() error
}

func (baseResponse BaseResponse) hasError() error {
	var err error
	switch baseResponse.Errno {
		case 0:
			err = nil
		case 20504:
			err = &DCEAPI.BalanceError{ErrCode:20504, ErrMsg:baseResponse.ErrMsg}
		case 20501, 20502:
			err = &DCEAPI.SymbolError{ErrCode:20501, ErrMsg:baseResponse.ErrMsg}
		case 20506, 20522, 20521:
			err = &DCEAPI.InvalidSignatureError{ErrCode:baseResponse.Errno, ErrMsg:baseResponse.ErrMsg}
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

func (aofex *Aofex) buildRequestBody(body map[string]string) (string, error) {
	formData := url.Values{}
	for k, v := range body {
		formData.Add(k, v)
	}
	return formData.Encode(), nil
}

func (aofex *Aofex) request(method, path string, params, body, headers map[string]string, model aofexResponse) error {
	var err error
	path = DCEAPI.BuildRequestUrl(path, params)
	bodyStr, err := aofex.buildRequestBody(body)
	if err != nil {
		return err
	}
	if headers != nil {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	res, err := DCEAPI.HttpRequest(method, path, bodyStr, headers)
	// fmt.Println(string(res))
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(res, model)
	if err != nil {
		return err
	}
	err = model.hasError()
	return err
}

func (aofex *Aofex) FetchBalance() ([]DCEAPI.Balance, error) {
	type balanceResponse struct {
		BaseResponse
		Result []DCEAPI.Balance
	}
	balance := &balanceResponse{}
	params := map[string]string{"show_all":"1"}
	headers := aofex.generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, params)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/wallet/list", params, nil, headers, balance)
	return balance.Result, err
}

func (aofex *Aofex) FetchMarkets() ([]DCEAPI.Market, error) {
	type MarketResponse struct {
		BaseResponse
		Result []DCEAPI.Market
	}
	market := &MarketResponse{}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/symbols", nil, nil, nil, market)
	return market.Result, err
}

func (aofex *Aofex) FetchTrades(symbol string) (DCEAPI.Trade, error) {
	type TradeResponse struct {
		BaseResponse
		Result DCEAPI.Trade
	}
	trade := &TradeResponse{}
	params := map[string]string{"symbol":symbol}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/trade", params, nil, nil, trade)
	return trade.Result, err
}

func (aofex *Aofex) FetchDepth(symbol string) (DCEAPI.OrderBook, error) {
	type DepthResponse struct {
		BaseResponse
		Result DCEAPI.OrderBook
	}
	depthResponse := &DepthResponse{}
	params := map[string]string{"symbol":symbol}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/depth", params, nil, nil, depthResponse)
	return depthResponse.Result, err
}

func (aofex *Aofex) FetchOHLCV(symbol, period string, size int) (DCEAPI.Kline, error) {
	type OHLCVResponse struct {
		BaseResponse
		Result DCEAPI.Kline
	}
	kline := &OHLCVResponse{}
	params := map[string]string{"symbol":symbol, "period":period, "size":strconv.Itoa(size)}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/kline", params, nil, nil, kline)
	return kline.Result, err
}

func (aofex *Aofex) FetchPercision(symbols ...string) (DCEAPI.Precision, error) {
	type PrecisionResponse struct {
		BaseResponse
		Result DCEAPI.Precision
	}
	precision := &PrecisionResponse{}
	params := map[string]string{}
	if len(symbols) != 0 {
		params["symbol"] = symbols[0]
	}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/precision", params, nil, nil, precision)
	return precision.Result, err 
}

func (aofex *Aofex) FetchKline24H(symbols ...string) ([]DCEAPI.Kline24H, error) {
	type Kline24HResponse struct {
		BaseResponse
		Result []DCEAPI.Kline24H
	}
	kline24HResponse := &Kline24HResponse{}
	params := map[string]string{}
	if len(symbols) != 0 {
		params["symbol"] = symbols[0]
	}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/24kline", params, nil, nil, kline24HResponse)
	return kline24HResponse.Result, err 
}

// 返回order
func (aofex *Aofex) Order(side, symbol, amount, price string) (DCEAPI.Order, error) {
	type OrderResponse struct {
		BaseResponse
		Result DCEAPI.Order
	}
	orderResponse := &OrderResponse{}
	body := map[string]string{
		"symbol": symbol,
		"type": side,
		"amount": amount,
		"price": price,
	}
	headers := aofex.generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, body)
	err := aofex.request("POST", aofex.SpotPath+"/openApi/entrust/add", nil, body, headers, orderResponse)
	return orderResponse.Result, err
}

func (aofex *Aofex) LimitBuyOrder(symbol, amount, price string) (DCEAPI.Order, error) {
	return aofex.Order("buy-limit", symbol, amount, price)
}

func (aofex *Aofex) LimitSellOrder(symbol, amount, price string) (DCEAPI.Order, error) {
	return aofex.Order("sell-limit", symbol, amount, price)
}

func (aofex *Aofex) MarketSellOrder(symbol, amount, price string) (DCEAPI.Order, error) {
	return aofex.Order("sell-market", symbol, amount, price)
}

func (aofex *Aofex) MarketBuyOrder(symbol, amount, price string) (DCEAPI.Order, error) {
	return aofex.Order("buy-market", symbol, amount, price)
}

// 使用order_sn来撤单
func (aofex *Aofex) CancelOrderByIDs(orderIDs ...string) ([]DCEAPI.Order, []DCEAPI.Order, error){
	var orderIDStr string
	for i, orderID := range orderIDs {
		orderIDStr += orderID
		if i != len(orderIDs)-1 {
			orderIDStr += ","
		}
	} 
	body := map[string]string{
		"order_ids": orderIDStr,
	}
	return aofex.cancelOrder(body)
}

// 使用symbol撤单
func (aofex *Aofex) CancelOrderBySymbol(symbol string) ([]DCEAPI.Order, []DCEAPI.Order, error) {
	body := map[string]string{
		"symbol": symbol,
	}
	return aofex.cancelOrder(body)
}

// 第一个返回值为成功订单 第二个返回值为失败订单
func (aofex *Aofex) cancelOrder(body map[string]string) ([]DCEAPI.Order, []DCEAPI.Order, error) {
	type CancelOrderResponse struct {
		BaseResponse
		Result struct{
			Success []string
			Failed []string
		}
	}
	cancelOrderResponse := &CancelOrderResponse{}
	headers := aofex.generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, body)
	err := aofex.request("POST", aofex.SpotPath+"/openApi/entrust/cancel", nil, body, headers, cancelOrderResponse)
	successOrders := []DCEAPI.Order{}
	failedOrders := []DCEAPI.Order{}
	for _, successOrder := range cancelOrderResponse.Result.Success {
		successOrders = append(successOrders, DCEAPI.Order{Order_sn:successOrder,})
	}
	for _, failedOrder := range cancelOrderResponse.Result.Failed {
		failedOrders = append(failedOrders, DCEAPI.Order{Order_sn:failedOrder,})
	}
	return successOrders, failedOrders, err
}

func (aofex *Aofex) FetchOpenOrders() ([]DCEAPI.Order, error) {
	type OpenOrderResponse struct {
		BaseResponse
		Result []DCEAPI.Order
	}
	openOrderResponse := &OpenOrderResponse{}
	headers := aofex.generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, nil)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/entrust/currentList", nil, nil, headers, openOrderResponse)
	return openOrderResponse.Result, err 
}

func (aofex *Aofex) FetchClosedOrder() ([]DCEAPI.Order, error) {
	type OpenOrderResponse struct {
		BaseResponse
		Result []DCEAPI.Order
	}
	openOrderResponse := &OpenOrderResponse{}
	headers := aofex.generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, nil)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/entrust/historyList", nil, nil, headers, openOrderResponse)
	return openOrderResponse.Result, err 
}

func (aofex *Aofex) FetchOrder(OrderID string) (DCEAPI.Order, error) {
	type OrderResponse struct {
		BaseResponse
		Result DCEAPI.Order
	}
	params := map[string]string{
		"order_sn": OrderID,
	}
	orderResponse := &OrderResponse{}
	headers := aofex.generateHeader(aofex.Exchange.Apikey, aofex.Exchange.Secret, params)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/entrust/status", params, nil, headers, orderResponse)
	return orderResponse.Result, err
}