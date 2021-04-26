package exchanges

import (
	"strconv"

	"github.com/PythonohtyP1900/DCEAPI"
)

func (aofex *Aofex) FetchBalance() ([]DCEAPI.Balance, error) {
	type balanceResponse struct {
		AofexBaseResponse
		Result []DCEAPI.Balance
	}
	balance := &balanceResponse{}
	params := map[string]string{"show_all":"1"}
	headers := aofex.generateHeader(aofex.Exchange.ApiKey, aofex.Exchange.Secret, params)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/wallet/list", params, nil, headers, balance)
	return balance.Result, err
}

func (aofex *Aofex) FetchMarkets() ([]DCEAPI.Market, error) {
	type MarketResponse struct {
		AofexBaseResponse
		Result []struct {
			Symbol string
			Base_currency string
			Quote_currency string
			Min_size float64
			Max_size float64
			Max_price float64
			Min_price float64
			Maker_fee float64
			Taker_fee float64
		}
	}
	response := &MarketResponse{}
	markets := []DCEAPI.Market{}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/symbols", nil, nil, nil, response)
	for _, result := range response.Result {
		market := DCEAPI.Market{
			Symbol: result.Symbol,
			BaseCurrency: result.Base_currency,
			QuoteCurrency: result.Quote_currency,
			MinSize: result.Min_size,
			MaxSize: result.Max_size,
			MinPrice: result.Min_price,
			MaxPrice: result.Max_price,
			MakerFee: result.Maker_fee,
			TakerFee: result.Taker_fee,
		}
		markets = append(markets, market)
	}
	return markets, err
}

func (aofex *Aofex) FetchTrades(symbol string) (DCEAPI.Trade, error) {
	type TradeResponse struct {
		AofexBaseResponse
		Result DCEAPI.Trade
	}
	trade := &TradeResponse{}
	params := map[string]string{"symbol":symbol}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/trade", params, nil, nil, trade)
	return trade.Result, err
}

func (aofex *Aofex) FetchDepth(symbol string) (DCEAPI.OrderBook, error) {
	type DepthResponse struct {
		AofexBaseResponse
		Result DCEAPI.OrderBook
	}
	depthResponse := &DepthResponse{}
	params := map[string]string{"symbol":symbol}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/depth", params, nil, nil, depthResponse)
	return depthResponse.Result, err
}

func (aofex *Aofex) FetchOHLCV(symbol, period string, size int) (DCEAPI.Kline, error) {
	type OHLCVResponse struct {
		AofexBaseResponse
		Result DCEAPI.Kline
	}
	kline := &OHLCVResponse{}
	params := map[string]string{"symbol":symbol, "period":period, "size":strconv.Itoa(size)}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/kline", params, nil, nil, kline)
	return kline.Result, err
}

func (aofex *Aofex) FetchPercision(symbols ...string) (DCEAPI.Precision, error) {
	type PrecisionResponse struct {
		AofexBaseResponse
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
		AofexBaseResponse
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
		AofexBaseResponse
		Result DCEAPI.Order
	}
	orderResponse := &OrderResponse{}
	body := map[string]string{
		"symbol": symbol,
		"type": side,
		"amount": amount,
		"price": price,
	}
	headers := aofex.generateHeader(aofex.Exchange.ApiKey, aofex.Exchange.Secret, body)
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
		AofexBaseResponse
		Result struct{
			Success []string
			Failed []string
		}
	}
	cancelOrderResponse := &CancelOrderResponse{}
	headers := aofex.generateHeader(aofex.Exchange.ApiKey, aofex.Exchange.Secret, body)
	err := aofex.request("POST", aofex.SpotPath+"/openApi/entrust/cancel", nil, body, headers, cancelOrderResponse)
	successOrders := []DCEAPI.Order{}
	failedOrders := []DCEAPI.Order{}
	for _, successOrder := range cancelOrderResponse.Result.Success {
		successOrders = append(successOrders, DCEAPI.Order{OrderID:successOrder,})
	}
	for _, failedOrder := range cancelOrderResponse.Result.Failed {
		failedOrders = append(failedOrders, DCEAPI.Order{OrderID:failedOrder,})
	}
	return successOrders, failedOrders, err
}

func (aofex *Aofex) FetchOpenOrders() ([]DCEAPI.Order, error) {
	type OpenOrderResponse struct {
		AofexBaseResponse
		Result []DCEAPI.Order
	}
	openOrderResponse := &OpenOrderResponse{}
	headers := aofex.generateHeader(aofex.Exchange.ApiKey, aofex.Exchange.Secret, nil)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/entrust/currentList", nil, nil, headers, openOrderResponse)
	return openOrderResponse.Result, err 
}

func (aofex *Aofex) FetchClosedOrder() ([]DCEAPI.Order, error) {
	type OpenOrderResponse struct {
		AofexBaseResponse
		Result []DCEAPI.Order
	}
	openOrderResponse := &OpenOrderResponse{}
	headers := aofex.generateHeader(aofex.Exchange.ApiKey, aofex.Exchange.Secret, nil)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/entrust/historyList", nil, nil, headers, openOrderResponse)
	return openOrderResponse.Result, err 
}

func (aofex *Aofex) FetchOrder(OrderID string) (DCEAPI.Order, error) {
	type OrderResponse struct {
		AofexBaseResponse
		Result DCEAPI.Order
	}
	params := map[string]string{
		"order_sn": OrderID,
	}
	orderResponse := &OrderResponse{}
	headers := aofex.generateHeader(aofex.Exchange.ApiKey, aofex.Exchange.Secret, params)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/entrust/status", params, nil, headers, orderResponse)
	return orderResponse.Result, err
}