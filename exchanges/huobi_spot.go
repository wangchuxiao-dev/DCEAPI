package exchanges

import (
	"github.com/PythonohtyP1900/DCEAPI"
	"github.com/shopspring/decimal"

	"strconv"
	"strings"
	"fmt"
)

func (huobi Huobi) FetchBalance() ([]DCEAPI.Balance, error) {
	url := fmt.Sprintf("/v1/account/accounts/%d/balance", huobi.SpotID)
	type BalanceResponse struct{
		huobiBaseResponse
		Data struct{
			Id int `json:"id"`
			Type string `json:"type"`
			List []struct{
				Currency string `json:"currency"`
				Type string `json:"type"`
				Balance decimal.Decimal `json:"balance"`
			} `json:"list"`
		}
	}
	balanceResponse := &BalanceResponse{}
	signMap := huobi.generateSignature(huobi.Path, "GET", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey, nil)
	err := huobi.request("GET", huobi.Path+url, signMap, nil, balanceResponse)
	var balances []DCEAPI.Balance
	for i, v := range balanceResponse.Data.List {
		if i % 2 == 0 {
			balance := DCEAPI.Balance{
				Currency: strings.ToUpper(v.Currency),
				Free: v.Balance,
				Frozen: balanceResponse.Data.List[i+1].Balance,
			}
			balances = append(balances, balance)
		}

	}
	return balances, err
}

func (huobi Huobi) FetchMarkets() ([]DCEAPI.Market, error) {
	url := "/v1/common/symbols"
	type MarketResponse struct {
		huobiBaseResponse
		Data []struct{
			Symbol string `json:"symbol"`
			BaseCurrency string `json:"base-currency"`
			QuoteCurrency string `json:"quote-currency"`
			PricePrecision int `json:"price-precision"`
			AmountPrecision int `json:"amount-precision"`
			LimitMinOrderAmount float64 `json:"limit-order-min-order-amt"`
			LimitMaxOrderAmount float64 `json:"limit-order-max-order-amt"`
			SellMarketMaxOrderAmount float64 `json:"sell-market-max-order-amt"`
			SellMarketMinOrderAmount float64 `json:"sell-market-min-order-amt"`
			BuyMarketMaxValue float64 `json:"buy-market-max-order-value"`
			MinOrderValue float64 `json:"min-order-value"`
		} `json:"data"`
	}
	marketResponse := &MarketResponse{}
	err := huobi.request("GET", huobi.Path+url, nil, nil, marketResponse)
	var markets []DCEAPI.Market 
	for _, v := range marketResponse.Data {
		market := DCEAPI.Market{
			Symbol: v.Symbol,
			BaseCurrency: v.BaseCurrency,
			QuoteCurrency: v.QuoteCurrency,
			PricePrecision: v.PricePrecision,
			AmountPrecision: v.AmountPrecision,
			LimitMinOrderAmount: v.LimitMinOrderAmount,
			LimitMaxOrderAmount: v.LimitMaxOrderAmount,
			SellMarketMaxOrderAmount: v.SellMarketMaxOrderAmount,
			SellMarketMinOrderAmount: v.SellMarketMinOrderAmount,
			BuyMarketMaxValue: v.BuyMarketMaxValue,
			MinOrderValue: v.MinOrderValue,
		}
		markets = append(markets, market)
	}
	return markets, err
}

func (huobi Huobi) Order(side, symbol, amount, price string) (DCEAPI.Order, error) {
	url := "/v1/order/orders/place"
	type OrderResponse struct {
		huobiBaseResponse
		OrderID string `json:"data"`
	}
	orderResponse := &OrderResponse{}
	body := map[string]interface{}{
		"account-id": strconv.Itoa(huobi.SpotID),
		"symbol": huobi.symbolFormatConversion(symbol),
		"type": side,
		"source": "spot-api",
		"price": price,
		"amount": amount,
	}
	signMap := huobi.generateSignature(huobi.Path, "POST", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey, nil)
	err := huobi.request("POST", huobi.Path+url, signMap, body, orderResponse)
	return DCEAPI.Order{OrderID:orderResponse.OrderID,}, err
}

func (huobi Huobi) LimitBuyOrder(symbol, amount, price string) (DCEAPI.Order, error){
	return huobi.Order("buy-limit", symbol, amount, price)
}

func (huobi Huobi) LimitSellOrder(symbol, amount, price string) (DCEAPI.Order, error){
	return huobi.Order("sell-limit", symbol, amount, price)
}

func (huobi Huobi) MarketBuyOrder(symbol, amount string) (DCEAPI.Order, error){
	return huobi.Order("buy-market", symbol, amount, "")
}

func (huobi Huobi) MarketSellOrder(symbol, amount string) (DCEAPI.Order, error){
	return huobi.Order("sell-market", symbol, amount, "")
}

type huobiOrder struct{
	OrderID int `json:"id"`
	Symbol string `json:"symbol"`
	Account int `json:"account-id"`
	Amount decimal.Decimal `json:"amount"`
	Price decimal.Decimal `json:"price"`
	CreateTime int `json:"created-at"`
	ClosedTime int `json:"finished-at"`
	Type string `json:"type"`
	Status string `json:"state"`
	Fee decimal.Decimal `json:"field-fees"`
	FilledAmountBase decimal.Decimal `json:"field-amount"`
	FilledAmountQuote decimal.Decimal `json:"field-cash-amount"`
} 

func (huobi Huobi) FetchOrder(orderID string) (DCEAPI.Order, error) {
	url := "/v1/order/orders/" + orderID
	type OrderRespose struct {
		huobiBaseResponse
		Data huobiOrder `json:"data"`
	}
	orderResponse := &OrderRespose{}
	signMap := huobi.generateSignature(huobi.Path, "GET", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey, nil)
	err := huobi.request("GET", huobi.Path+url, signMap, nil, orderResponse)
	if err != nil{
		return DCEAPI.Order{}, err
	}
	order := huobi.orderFormatConversionToDCEFormat(orderResponse.Data)
	return *order, err
}

func (huobi Huobi) CancelOrder(orderID string) (error) {
	url := fmt.Sprintf("/v1/order/orders/%s/submitcancel", orderID)
	type CancelResponse struct {
		huobiBaseResponse
		Data string `json:"data"`
	}
	cancelResponse := &CancelResponse{}
	signMap := huobi.generateSignature(huobi.Path, "POST", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey, nil)
	err := huobi.request("POST", huobi.Path+url, signMap, nil, cancelResponse)
	return err
}

func (huobi Huobi) BatchCancelOrder(orderIDs ...string) ([]DCEAPI.Order, []DCEAPI.Order, error) {
	url := "/v1/order/orders/batchcancel"
	body := map[string]interface{}{
		"order-ids":orderIDs,
	}
	type BatchCancelResponse struct{
		huobiBaseResponse
		Data struct{
			Success []string `json:"success"`
			Failed []struct{
				OrderID string `json:"order-id"`
			} `json:"failed"`
		} `json:"data"`
	}
	batchCancelResponse := &BatchCancelResponse{}
	signMap := huobi.generateSignature(huobi.Path, "POST", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey, nil)
	err := huobi.request("POST", huobi.Path+url, signMap, body, batchCancelResponse)
	successOrders := []DCEAPI.Order{}
	failedOrders := []DCEAPI.Order{}
	for _, v := range batchCancelResponse.Data.Success {
		successOrders = append(successOrders, DCEAPI.Order{
			OrderID: v,
		})
	}
	for _, v := range batchCancelResponse.Data.Failed {
		failedOrders = append(failedOrders, DCEAPI.Order{
			OrderID: v.OrderID,
		})
	}
	return successOrders, failedOrders, err
}

func (huobi Huobi) CancelOrderBySymbol(symbols ...string) ([]DCEAPI.Order, []DCEAPI.Order, error){
	return []DCEAPI.Order{}, []DCEAPI.Order{}, DCEAPI.UnsupportMethodError{MethodName:"CancelOrderBySymbol", ExchangeName:"huobipro"}
}

func (huobi Huobi) FetchOpenOrders(symbol ...string) ([]DCEAPI.Order, error) {
	return []DCEAPI.Order{}, DCEAPI.UnsupportMethodError{MethodName:"FetchOpenOrders", ExchangeName:"huobipro"}
}

func (huobi Huobi) FetchClosedOrders(symbol ...string) ([]DCEAPI.Order, error) {
	return []DCEAPI.Order{}, DCEAPI.UnsupportMethodError{MethodName:"FetchClosedOrders", ExchangeName:"huobipro"}
}

// 返回指定交易对最新的一个交易记录
func (huobi Huobi) FetchTrade(symbol string) ([]DCEAPI.Trade, error) {
	url := "/market/trade"
	params := map[string]string{
		"symbol": huobi.symbolFormatConversion(symbol),
	}
	type TradeResponse struct {
		huobiBaseResponse
		Tick struct{
			Data []struct{
				Ts int `json:"ts"`
				Price decimal.Decimal `json:"price"`
				Amount decimal.Decimal `json:"amount"`
				Side string `json:"direction"`
			} `json:"data"`
		} `json:"tick"`
	}
	tradeResponse := &TradeResponse{}
	err := huobi.request("GET", huobi.Path+url, params, nil, tradeResponse)
	trades := []DCEAPI.Trade{}
	fmt.Println(tradeResponse)
	for _, v := range tradeResponse.Tick.Data {
		trade := DCEAPI.Trade{
			Amount: v.Amount,
			Price: v.Price,
			Side: v.Side,
			Ts: v.Ts,
		}
		trades = append(trades, trade)
	}
	return trades, err
}

// 获取交易记录
func (huobi Huobi) FetchTrades(symbol string, size string) ([]DCEAPI.Trade, error) {
	url := "/market/history/trade"
	params := map[string]string{
		"symbol": huobi.symbolFormatConversion(symbol),
	}
	if size != "" {
		params["size"] = size
	}
	type TradeResponse struct {
		huobiBaseResponse
		Data []struct{
			Data []struct{
				Amount decimal.Decimal `json:"amount"`
				Ts int `json:"ts"`
				Price decimal.Decimal `json:"price"`
				Side string `json:"direction"`
			} `json:"data"`
		} `json:"data"`
	}
	tradeResponse := &TradeResponse{}
	err := huobi.request("GET", huobi.Path+url, params, nil, tradeResponse)
	var trades []DCEAPI.Trade
	for _, v := range tradeResponse.Data{
		for _, v2 := range v.Data {
			trade := DCEAPI.Trade{
				Symbol: symbol,
				Amount: v2.Amount,
				Price: v2.Price,
				Ts: v2.Ts,
				Side: v2.Side,
			}
			trades = append(trades, trade)
		}
	}
	return trades, err
}

func (huobi Huobi) FetchOrderBook(symbol string, params map[string]string) (DCEAPI.OrderBook, error) {
	url := "/market/depth"
	reqParams := map[string]string{
		"symbol": huobi.symbolFormatConversion(symbol),
	}
	for k, v := range params {
		reqParams[k] = v
	}
	type OrderBookResponse struct {
		huobiBaseResponse
		Tick struct {
			Bids [][2]decimal.Decimal `json:"bids"`
			Asks [][2]decimal.Decimal `json:"asks"`
			Ts int `json:"ts"`
		} `json:"tick"`
	}
	orderBookResponse := &OrderBookResponse{}
	err := huobi.request("GET", huobi.Path+url, reqParams, nil, orderBookResponse)
	orderBook := DCEAPI.OrderBook{
		Bids: orderBookResponse.Tick.Bids,
		Asks: orderBookResponse.Tick.Asks,
		Symbol: symbol,
		Ts: orderBookResponse.Tick.Ts,
	}
	return orderBook, err
}

func (huobi Huobi) FetchOHLCV(symbol, period, size string) ([]DCEAPI.Kline, error) {
	url := "/market/history/kline"
	reqParams := map[string]string{
		"symbol": huobi.symbolFormatConversion(symbol),
		"period": period,
	} 
	type OHLCVResponse struct{
		huobiBaseResponse
		Data []struct{
			Amount decimal.Decimal `json:"amount"`
			Count decimal.Decimal `json:"count"`
			Open decimal.Decimal `json:"open"`
			Close decimal.Decimal `json:"close"`
			Low decimal.Decimal `json:"low"`
			High decimal.Decimal `json:"high"`
			Vol decimal.Decimal `json:"vol"`
		} `json:"data"`
	}
	ohlcvResponse := &OHLCVResponse{}
	err := huobi.request("GET", huobi.Path+url, reqParams, nil, ohlcvResponse)
	klines := []DCEAPI.Kline{}
	for _, v := range ohlcvResponse.Data {
		klines = append(klines, DCEAPI.Kline{
			Amount: v.Amount,
			Count: v.Count,
			Open: v.Open,
			Close: v.Close,
			Low: v.Low,
			High: v.High,
			Vol: v.Vol,
		})
	}
	return klines, err
}

func (huobi Huobi) FetchOHLCV24H(symbol string) (DCEAPI.Kline, error) {
	url := "/market/detail"
	reqParams := map[string]string{"symbol":huobi.symbolFormatConversion(symbol)}
	type OHLCVResponse struct{
		huobiBaseResponse
		Tick struct{
			Amount decimal.Decimal `json:"amount"`
			Count decimal.Decimal `json:"count"`
			Open decimal.Decimal `json:"open"`
			Close decimal.Decimal `json:"close"`
			Low decimal.Decimal `json:"low"`
			High decimal.Decimal `json:"high"`
			Vol decimal.Decimal `json:"vol"`
		} `json:"Tick"`
	}
	ohlcvResponse := &OHLCVResponse{}
	err := huobi.request("GET", huobi.Path+url, reqParams, nil, ohlcvResponse)
	kline := DCEAPI.Kline{
		Amount: ohlcvResponse.Tick.Amount,
		Count: ohlcvResponse.Tick.Count,
		Open: ohlcvResponse.Tick.Open,
		Close: ohlcvResponse.Tick.Close,
		Low: ohlcvResponse.Tick.Low,
		High: ohlcvResponse.Tick.High,
		Vol: ohlcvResponse.Tick.Vol,
	}
	return kline, err
}	

func (huobi Huobi) FetchFee(symbols... string) (error) {
	url := "/v2/reference/transact-fee-rate"
	type FeeResponse struct {
		huobiBaseResponse
		Data []struct {
			Symbol string `json:"symbol"`
			MakerFee string `json:"makerFeeRate"`
			TakerFee string `json:"takerFeeRate"`
		} `json:"data"`
	}
	feeResponse := &FeeResponse{}
	signMap := huobi.generateSignature(huobi.Path, "GET", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey, map[string]string{"symbol":"btcusdt, ethusdt"})
	err := huobi.request("GET", huobi.Path+url, signMap, nil, feeResponse)
	fmt.Println(feeResponse.Data)
	return err
}

func (huobi Huobi) orderFormatConversionToDCEFormat(horder huobiOrder) (*DCEAPI.Order){
	status := horder.Status
	switch status {
		case "filled":
			status = "closed"
		case "submitted":
			status = "open"
		case "canceled":
			status = "canceled"
	}
	temp := strings.Split(horder.Type, "-")
	order := DCEAPI.Order{
		OrderID: strconv.Itoa(horder.OrderID),
		Symbol: huobi.symbolFormatConversionToDCEFormat(horder.Symbol),
		Type: temp[1],
		Side: temp[0],
		CreateTime: horder.CreateTime/1000,
		ClosedTime: horder.ClosedTime/1000,
		Price: horder.Price,
		Amount: horder.Amount,
		DealPrice: horder.Price,
		DealAmountBase: horder.FilledAmountBase,
		DealAmountQuote: horder.FilledAmountQuote,
		Fee: horder.Fee,
		Status: status,
	}
	return &order
}