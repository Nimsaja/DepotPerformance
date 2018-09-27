package yahoo

// QuoteResponse ...
type QuoteResponse struct {
	QR Output `json:"quoteResponse"`
}

// Output ...
type Output struct {
	Res   []Result    `json:"result"`
	Error interface{} `json:"error"`
}

// Result ...
type Result struct {
	Cur    string  `json:"currency"`
	Name   string  `json:"longName"`
	Close  float32 `json:"regularMarketPreviousClose"`
	Price  float32 `json:"regularMarketPrice"`
	Symbol string  `json:"symbol"`
	Time   int64   `json:"regularMarketTime"`
}
