package yahoo

import (
	"testing"

	"github.com/Nimsaja/DepotPerformance/depot"
)

func TestConvertJSON2ResultIfResultArrayIsEmpty(t *testing.T) {
	b := []byte(`{"quoteResponse":{"result":[],"error":null}}`)

	_, err := convertJSON2Result(b)
	if err == nil {
		t.Errorf("No error expected")
	}
}

func TestConvertJSON2ResultIfStockIsFound(t *testing.T) {
	b := []byte(`{"quoteResponse":{"result":
		[{"currency":"USD",
		"regularMarketPrice":1199.89,
		"regularMarketPreviousClose":1180.49,
		"longName":"Alphabet Inc.",
		"shortName":"Alphabet Inc.",
		"symbol":"GOOG"}],
		"error":null}}`)

	r, err := convertJSON2Result(b)
	if err != nil {
		t.Errorf("No error expected")
	}

	if r.Price != 1199.89 {
		t.Errorf("Expected %v, got %v", 1199.89, r.Price)
	}
}

func TestGet(t *testing.T) {
	s := depot.Stock{Name: "Google", Symbol: "ABEC.DE"}

	r := Get(s)

	if r.Name != "Alphabet Inc." {
		t.Errorf("Name of stock should be %v. Got %v.", "Alphabet Inc.", r.Name)
	}
	if r.Cur != "EUR" {
		t.Errorf("Currency of stock should be %v. Got %v.", "EUR", r.Cur)
	}
}
