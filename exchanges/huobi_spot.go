package exchanges

import (
	"github.com/PythonohtyP1900/DCEAPI"

	"strconv"
)

func (huobi Huobi) Order(side, symbol, amount, price string) (DCEAPI.Order, error) {
	url := "/v1/order/orders/place"
	type OrderResponse struct {
		huobiBaseResponse
		OrderID string `json:"data"`
	}
	orderResponse := &OrderResponse{}
	body := map[string]string{
		"account-id": strconv.Itoa(huobi.SpotID),
		"symbol": symbol,
		"type": side,
		"source": "spot-api",
		"price": price,
		"amount": amount,
	}
	signMap := huobi.generateSignature(huobi.Path, "POST", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey)
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

func (huobi Huobi) FetchOrder(orderID string) (DCEAPI.Order, error){
	body := map[string]string{
		"order-id": orderID,
	}
	type OrderRespose struct {
		huobiBaseResponse
		Data struct {
			Symbol string `json:"symbol"`

		}
	}
	orderRespnse := &OrderRespose{}
	signMap := huobi.generateSignature(huobi.Path, "GET", url, huobi.Exchange.Secret, huobi.Exchange.ApiKey)
	err := huobi.request("GET", huobi.Path+url, signMap, body, orderResponse)
}
