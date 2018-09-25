package depot

import (
	"math"
)

//Stock structure
type Stock struct {
	Name    string
	Symbol  string
	Count   float32
	powEuro int8
}

//Get gets the portfolio
func Get() []Stock {
	return []Stock{
		Stock{Name: "Google", Symbol: "GOOG", Count: 0.211, powEuro: 1},
		Stock{Name: "Amazon", Symbol: "AMZN", Count: 0.056, powEuro: 1},
		Stock{Name: "Netflix", Symbol: "NFLX", Count: 2, powEuro: 1},
		Stock{Name: "Siemens", Symbol: "SIE-D", Count: 5, powEuro: 0},
	}
}

//AsEuro converts dollar values to euro if necessary
func (s Stock) AsEuro(v float32, euro float32) float32 {
	return v * float32(math.Pow(float64(euro), float64(s.powEuro)))
}
