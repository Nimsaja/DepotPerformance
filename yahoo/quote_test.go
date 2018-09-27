package yahoo

import "testing"

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
