package gvar

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
