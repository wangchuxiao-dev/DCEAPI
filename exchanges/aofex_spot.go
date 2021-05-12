package exchanges

import (
	"github.com/PythonohtyP1900/DCEAPI"
	"github.com/shopspring/decimal"

	"strconv"
)

func (aofex *Aofex) FetchBalance() ([]DCEAPI.Balance, error) {
	type BalanceResponse struct {
		AofexBaseResponse
		Data []struct{
			Currency string `json:"currency"`
			Free decimal.Decimal `json:"Available"`
			Frozen decimal.Decimal `json:"Frozen"`
		} `json:"Result"`
	}
	balanceResponse := &BalanceResponse{}
	params := map[string]string{"show_all":"1"}
	headers := aofex.generateHeader(aofex.Exchange.ApiKey, aofex.Exchange.Secret, params)
	err := aofex.request("GET", aofex.SpotPath+"/openApi/wallet/list", params, nil, headers, balanceResponse)
	var balances []DCEAPI.Balance
	for _, v := range balanceResponse.Data {
		balances = append(balances, DCEAPI.Balance{
			Currency: v.Currency,
			Free: v.Free,
			Frozen: v.Frozen,
		})
	}
	return balances, err
}

func (aofex *Aofex) FetchMarkets() ([]DCEAPI.Market, error) {
	markets := []DCEAPI.Market{}
	type MarketResponse struct {
		AofexBaseResponse
		Result []struct {
			Symbol string
			BaseCurrency string `json:"base_currency"`
			QuoteCurrency string `json:"quote_currency"`
			MinSize float64 `json:"min_size"`
			MaxSize float64 `json:"max_size"`
			MaxPrice float64 `json:"max_price"`
			MinPrice float64 `json:"min_price"`
			MakerFee float64 `json:"maker_fee"`
			TakerFee float64 `json:"taker_fee"`
		} `json:"result"`
	}
	response := &MarketResponse{}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/symbols", nil, nil, nil, response)
	if err != nil {
		return markets, err
	}
	type PrecisionStruct struct{
		AofexBaseResponse
		Result map[string]map[string]string	
	}
	precisionResponse := &PrecisionStruct{}
	err = aofex.request("GET", aofex.SpotPath+"/openApi/market/precision", nil, nil, nil, precisionResponse)
	for _, result := range response.Result {
		for symbol, v:= range precisionResponse.Result{
			if result.Symbol == symbol {
				pricePrecision, _ := strconv.Atoi(v["price"])
				amountPrecision, _ := strconv.Atoi(v["amount"])
				minQuantity, _ := strconv.ParseFloat(v["minQuantity"], 64)
				maxQuantity, _ := strconv.ParseFloat(v["maxQuantity"], 64)
				market := DCEAPI.Market{
					Symbol: result.Symbol,
					BaseCurrency: result.BaseCurrency,
					QuoteCurrency: result.QuoteCurrency,
					PricePrecision: pricePrecision,
					AmountPrecision: amountPrecision,
					LimitMinOrderAmount: minQuantity,
					SellMarketMinOrderAmount: minQuantity,
					SellMarketMaxOrderAmount: maxQuantity,
					LimitMaxOrderAmount: maxQuantity,
				}
				markets = append(markets, market)
			}	
		}
	}
	return markets, err
}

func (aofex *Aofex) FetchTrade(symbol string) ([]DCEAPI.Trade, error) {
	return []DCEAPI.Trade{}, DCEAPI.UnsupportMethodError{MethodName:"FetchTrade", ExchangeName:"AOFEX"}
}

func (aofex *Aofex) FetchTrades(symbol, size string) ([]DCEAPI.Trade, error) {
	url := "/openApi/market/trade"
	type TradeResponse struct {
		AofexBaseResponse
		Result struct{
			Data []struct{
				Ts int `json:"ts"`
				Amount decimal.Decimal `json:"amount"`
				Price decimal.Decimal `json:"price"`
				Side string `json:"direction"`
			} `json:"data"`
		} `json:"result"`
	}
	trade := &TradeResponse{}
	params := map[string]string{"symbol":aofex.SymbolFormatConversion(symbol)}
	err := aofex.request("GET", aofex.SpotPath+url, params, nil, nil, trade)
	trades := []DCEAPI.Trade{}
	for _, v := range trade.Result.Data {
		trade := DCEAPI.Trade{
			Symbol: symbol,
			Amount: v.Amount,
			Price: v.Price,
			Side: v.Side,
			Ts: v.Ts,
		}
		trades = append(trades, trade)
	}
	return trades, err
}

func (aofex *Aofex) FetchOrderBook(symbol string, params map[string]string) (DCEAPI.OrderBook, error) {
	url := "/openApi/market/depth"
	type DepthResponse struct {
		AofexBaseResponse
		Result struct{
			Symbol string `json:"symbol"`
			Ts int `json:"ts"`
			Bids [][2]decimal.Decimal `json:"bids"`
			Asks [][2]decimal.Decimal `json:"asks"`
		} `json:"result"`
	}
	depthResponse := &DepthResponse{}
	reqParams := map[string]string{"symbol":aofex.SymbolFormatConversion(symbol)}
	for k, v := range params {
		reqParams[k] = v
	}
	err := aofex.request("GET", aofex.SpotPath+url, reqParams, nil, nil, depthResponse)
	var orderBook DCEAPI.OrderBook
	orderBook.Ts = depthResponse.Result.Ts
	orderBook.Symbol = symbol
	orderBook.Bids = depthResponse.Result.Bids
	orderBook.Asks = depthResponse.Result.Asks
	return orderBook, err
}

func (aofex *Aofex) FetchOHLCV(symbol, period, size string) ([]DCEAPI.Kline, error) {
	type OHLCVResponse struct {
		AofexBaseResponse
		Result struct{
			Data []struct{
				Amount decimal.Decimal `json:"amount"`
				Count decimal.Decimal `json:"count"`
				Open decimal.Decimal `json:"open"`
				Close decimal.Decimal `json:"close"`
				Low decimal.Decimal `json:"low"`
				High decimal.Decimal `json:"high"`
				Vol decimal.Decimal `json:"vol"`
			} `json:"data"`
		} `json:"result"`
	}
	klinesResponse := &OHLCVResponse{}
	params := map[string]string{"symbol":aofex.SymbolFormatConversion(symbol), "period":period, "size":size}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/kline", params, nil, nil, klinesResponse)
	kilnes := []DCEAPI.Kline{}
	for _, v := range klinesResponse.Result.Data{
		kilnes = append(kilnes, DCEAPI.Kline{
			Amount: v.Amount,
			Count: v.Count,
			Open: v.Open,
			Close: v.Close,
			Low: v.Low,
			High: v.High,
			Vol: v.Vol,
		})
	}
	return kilnes, err
}

func (aofex *Aofex) FetchOHLCV24H(symbol string) (DCEAPI.Kline, error) {
	type Kline24HResponse struct {
		AofexBaseResponse
		Result []struct{
			Data struct{
				Amount decimal.Decimal `json:"amount"`
				Count decimal.Decimal `json:"count"`
				Open decimal.Decimal `json:"open"`
				Close decimal.Decimal `json:"close"`
				Low decimal.Decimal `json:"low"`
				High decimal.Decimal `json:"high"`
				Vol decimal.Decimal `json:"vol"`
			} `json:"data"`
		} `json:"result"`
	}
	kline24HResponse := &Kline24HResponse{}
	params := map[string]string{"symbol": aofex.SymbolFormatConversion(symbol)}
	err := aofex.request("GET", aofex.SpotPath+"/openApi/market/24kline", params, nil, nil, kline24HResponse)
	kilne := DCEAPI.Kline{
		Amount: kline24HResponse.Result[0].Data.Amount,
		Count: kline24HResponse.Result[0].Data.Count,
		Open: kline24HResponse.Result[0].Data.Open,
		Close: kline24HResponse.Result[0].Data.Close,
		Low: kline24HResponse.Result[0].Data.Low,
		High: kline24HResponse.Result[0].Data.High,
		Vol: kline24HResponse.Result[0].Data.Vol,
	}
	
	return kilne, err 
}

// 返回order
func (aofex *Aofex) Order(side, symbol, amount, price string) (DCEAPI.Order, error) {
	type OrderResponse struct {
		AofexBaseResponse
		Result DCEAPI.Order
	}
	orderResponse := &OrderResponse{}
	body := map[string]string{
		"symbol": aofex.SymbolFormatConversion(symbol),
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

func (aofex *Aofex) MarketSellOrder(symbol, amount string) (DCEAPI.Order, error) {
	return aofex.Order("sell-market", symbol, amount, "")
}

func (aofex *Aofex) MarketBuyOrder(symbol, amount string) (DCEAPI.Order, error) {
	return aofex.Order("buy-market", symbol, amount, "")
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

// 单个order_id撤单
func (aofex *Aofex) CancelOrder(OrderID string) ([]DCEAPI.Order, []DCEAPI.Order, error) {
	body := map[string]string{
		"order_ids": OrderID,
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