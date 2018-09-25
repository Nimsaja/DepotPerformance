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
		Stock{Name: "XING", Symbol: "O1BC-D", Count: 2, powEuro: 0},
		Stock{Name: "Biotech", Symbol: "BBZA-D", Count: 3, powEuro: 0},
		Stock{Name: "Auto&Robotic", Symbol: "2B76", Count: 33, powEuro: 0},
		Stock{Name: "TecDax", Symbol: "TDXP", Count: 10, powEuro: 0},
		Stock{Name: "Oekoworld", Symbol: "OE7A", Count: 0.523, powEuro: 0},
	}
}

//AsEuro converts dollar values to euro if necessary
func (s Stock) AsEuro(v float32, euro float32) float32 {
	return v * float32(math.Pow(float64(euro), float64(s.powEuro)))
}
