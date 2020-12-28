package exchange

import (
	"crypto/sha1"
	"exchange"
)

const Path string = "https://aofex.com/"
const CNPath string = "https://aofex.co/"
const SpotPath string = "https://openapi.aofex.com/"
const SwapPath string = "https://openapi-contract.aofex.com/"
const DebugSpotPath string = ""
const DebugSwapPath string = ""

type Aofex struct {
	Path string
	SpotPath string
	SwapPath string
	Exchange *Exchange
}

func NewAofex(debug bool, secert, apikey string) *Aofex {
	var spotPath string
	var swapPath string
	if debug == true {
		spotPath = DebugSpotPath
		swapPath = DebugSwapPath
	} else {
		spotPath = SpotPath
		swapPath = SwapPath
	}
	aofex := &Aofex{
		Path: Path,
		SpotPath: spotPath,
		SwapPath: swapPath,
		Exchange: &Exchange{
			secert: secert,
			apikey: apikey,
			Debug: debug,
		},
	}
	return aofex
}

func sign(apikey string, token string) {

}

func getHeader() {
	
}

func (aofex *Aofex) FetchMarkets() (string, error) {
	route := "openApi/market/symbols"
	url := aofex.SpotPath + route
	res, err := BaseRequest("GET", url, "")
	return res, err
}

func (aofex *Aofex) FetchTrades() (string, error) {
	route := "openApi/market/trade"
	url := aofex.SpotPath + route
	res, err := BaseRequest("GET", url, "")
	return res, err
}

func (aofex *Aofex) LimitBuyOrder(symbol string, amount, price float64) {
	
}

func (aofex *Aofex) LimitSellOrder(symbol string, amount, price float64) {

}

func (aofex *Aofex) FetchBalance() {

}
