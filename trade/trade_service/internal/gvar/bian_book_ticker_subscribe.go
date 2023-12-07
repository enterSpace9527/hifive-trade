package gvar

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/models/trade_model"
	"time"
)

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
