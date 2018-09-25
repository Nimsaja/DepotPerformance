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
		Stock{Name: "Google", Symbol: "ABEC.DE", Count: 0.211},
		Stock{Name: "Amazon", Symbol: "AMZ.DE", Count: 0.056},
		Stock{Name: "Netflix", Symbol: "NFC.DE", Count: 2},
		Stock{Name: "Siemens", Symbol: "SIE.DE", Count: 5},
		Stock{Name: "XING", Symbol: "O1BC.SG", Count: 2},
		Stock{Name: "Biotech", Symbol: "DWWD.BE", Count: 3},
		Stock{Name: "Auto&Robotic", Symbol: "2B76.F", Count: 33},
		Stock{Name: "TecDax", Symbol: "EXS2.F", Count: 10},
		Stock{Name: "Oekoworld", Symbol: "OE7A.MU", Count: 0.523},
	}
}
