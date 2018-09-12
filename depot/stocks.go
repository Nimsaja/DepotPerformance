package depot

//Stock structure
type Stock struct {
	Name   string
	Symbol string
	Count  int
}

//Get gets the portfolio
func Get() []Stock {
	return []Stock{
		Stock{Name: "Google", Symbol: "GOOG", Count: 1},
		Stock{Name: "Amazon", Symbol: "AMZN", Count: 2},
	}
}
