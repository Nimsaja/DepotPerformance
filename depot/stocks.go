package depot

//Stock structure
type Stock struct {
	Name   string
	Symbol string
	Count  float32
	Buy    float32
}

//Get gets the portfolio
func Get() []Stock {
	return []Stock{
		Stock{Name: "Google", Symbol: "ABEC.DE", Count: 0.211, Buy: 1069.138},
		Stock{Name: "Amazon", Symbol: "AMZ.DE", Count: 0.056, Buy: 1776.515},
		Stock{Name: "Netflix", Symbol: "NFC.DE", Count: 2, Buy: 224.25},
		Stock{Name: "Siemens", Symbol: "SIE.DE", Count: 5, Buy: 106.02},
		Stock{Name: "XING", Symbol: "O1BC.F", Count: 2, Buy: 328.765},
		Stock{Name: "Biotech", Symbol: "DWWD.SG", Count: 3, Buy: 195.693},
		Stock{Name: "Auto&Robotic", Symbol: "2B76.F", Count: 33, Buy: 7.051},
		Stock{Name: "TecDax", Symbol: "EXS2.F", Count: 10, Buy: 25.01},
		Stock{Name: "Oekoworld", Symbol: "OE7A.SG", Count: 0.523, Buy: 191.051},
	}
}

//SumBuy gets the sum of spended money
func SumBuy() float32 {
	var sum float32
	for _, s := range Get() {
		sum += s.Buy * s.Count
	}
	return sum
}
