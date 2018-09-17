package depot

//Stock structure
type Stock struct {
	Name   string
	Symbol string
	Count  float32
}

//Get gets the portfolio
func Get() []Stock {
	return []Stock{
		Stock{Name: "Google", Symbol: "GOOG", Count: 0.211},
		Stock{Name: "Amazon", Symbol: "AMZN", Count: 0.056},
		Stock{Name: "Netflix", Symbol: "NFLX", Count: 2},
		Stock{Name: "Siemens", Symbol: "SIE-D", Count: 5},
	}
}
