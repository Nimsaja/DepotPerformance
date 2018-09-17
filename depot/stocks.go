package depot

//Stock structure
type Stock struct {
	Name     string
	Symbol   string
	Count    float32
	ConvEuro bool
}

//Get gets the portfolio
func Get() []Stock {
	return []Stock{
		Stock{Name: "Google", Symbol: "GOOG", Count: 0.211, ConvEuro: true},
		Stock{Name: "Amazon", Symbol: "AMZN", Count: 0.056, ConvEuro: true},
		Stock{Name: "Netflix", Symbol: "NFLX", Count: 2, ConvEuro: true},
		Stock{Name: "Siemens", Symbol: "SIE-D", Count: 5, ConvEuro: false},
	}
}

//ConvertToEuro ...
func ConvertToEuro(b bool, v float32, euro float32) float32 {
	if b {
		return v * euro
	}
	return v
}
