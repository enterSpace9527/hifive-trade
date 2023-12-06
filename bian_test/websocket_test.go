package bian_test

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"testing"
)

func TestWebsocket(t *testing.T) {

	binance.UseTestnet = true
	binance.WebsocketKeepalive = true

	//订阅用户数据变动
	doneC, _, err := binance.WsUserDataServe(apiKey, userDataHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	//订阅关注的市场数据变动
	//doneC, _, err = binance.WsAllBookTickerServe(bookTickerHandler, errHandler)
	doneC, _, err = binance.WsBookTickerServe("ETHUSDT", bookTickerHandler, errHandler)
	doneC, _, err = binance.WsBookTickerServe("BNBUSDT", bookTickerHandler, errHandler)
	fmt.Println("begin start")
	<-doneC
}

func bookTickerHandler(event *binance.WsBookTickerEvent) {
	fmt.Printf("bookTickerHandler event: %+v\n", event)
}

func userDataHandler(event *binance.WsUserDataEvent) {
	fmt.Printf("userDataHandler event: %+v\n", event)
}

func errHandler(err error) {
	fmt.Printf("errHandler err: %+v\n", err.Error())
	go TestWebsocket(nil)
}
