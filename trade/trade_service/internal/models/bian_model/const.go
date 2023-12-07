package bian_model

type TradeType string

const (
	TradeTypeLimit  TradeType = "LIMIT"
	TradeTypeMarket TradeType = "MARKET"
)

type TradeSide string

const (
	TradeSideBuy  TradeSide = "BUY"
	TradeSideSell TradeSide = "SELL"
)

type TradeStatus string

const (
	TradeStatusNew             TradeStatus = "NEW"
	TradeStatusPartiallyFilled TradeStatus = "PARTIALLY_FILLED"
	TradeStatusFilled          TradeStatus = "FILLED"
)
