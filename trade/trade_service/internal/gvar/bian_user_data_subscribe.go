package gvar

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/config"
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
