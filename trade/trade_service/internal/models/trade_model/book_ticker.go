package trade_model

import "time"

type BookTicker struct {
	Symbol       string
	BestBidPrice string
	BestBidQty   string
	BestAskPrice string
	BestAskQty   string
	UpdateTime   time.Time
}
