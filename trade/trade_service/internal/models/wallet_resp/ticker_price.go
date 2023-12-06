package wallet_resp

type TickerElement struct {
	S  string `json:"s"`
	P  string `json:"P"`
	C  string `json:"c"`
	P1 string `json:"p"`
	I  string `json:"i"`
}

type TickerPrice struct {
	List []TickerElement `json:"list"`
}
