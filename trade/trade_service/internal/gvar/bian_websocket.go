package gvar

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/config"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/models/trade_model"
	"time"
)

var gApikey = ""

func initBinanceUserDataSubscribe(c *config.Config) error {

	binance.WebsocketKeepalive = true
	if gApikey == "" {
		account1 := c.HifiveAccounts["Account1"]
		apiKey := account1["ApiKey"]
		gApikey = apiKey
	}

	if gApikey == "" {
		return fmt.Errorf("apikey is empty")
	}
	_, _, err := binance.WsUserDataServe(gApikey, userDataHandler, userDataErrorHandler)
	if err != nil {
		return err
	}
	return nil
}

func userDataHandler(event *binance.WsUserDataEvent) {
	fmt.Printf("userDataHandler event: %+v\n", event)
}

func userDataErrorHandler(err error) {
	fmt.Printf("userDataErrorHandler err: %+v\n", err.Error())
	for {
		time.Sleep(2 * time.Second)
		err := initBinanceUserDataSubscribe(nil)
		if err == nil {
			break
		}
	}
}

func initBinanceBookTickerSubscribe() error {

	binance.WebsocketKeepalive = true

	_, _, err := binance.WsBookTickerServe("BTCUSDT", bookTickerHandler, bookTickerErrorHandler)
	if err != nil {
		return err
	}

	_, _, err = binance.WsBookTickerServe("ETHUSDT", bookTickerHandler, bookTickerErrorHandler)
	if err != nil {
		return err
	}

	_, _, err = binance.WsBookTickerServe("BNBUSDT", bookTickerHandler, bookTickerErrorHandler)
	if err != nil {
		return err
	}
	return nil
}

func bookTickerHandler(event *binance.WsBookTickerEvent) {
	bookTicker := trade_model.BookTicker{}
	bookTicker.Symbol = event.Symbol
	bookTicker.BestAskQty = event.Symbol
	bookTicker.BestAskPrice = event.Symbol
	bookTicker.BestBidQty = event.Symbol
	bookTicker.BestBidPrice = event.Symbol
	bookTicker.UpdateTime = time.Now()
	bookTickerMap[event.Symbol] = bookTicker
}

func bookTickerErrorHandler(err error) {
	fmt.Printf("userDataErrorHandler err: %+v\n", err.Error())

	for {
		time.Sleep(2 * time.Second)
		err := initBinanceBookTickerSubscribe()
		if err == nil {
			break
		}
	}
}
