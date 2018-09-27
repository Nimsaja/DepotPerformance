package yahoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Nimsaja/DepotPerformance/depot"
)

var url = "https://query1.finance.yahoo.com/v7/finance/quote?lang=en-US&region=US&corsDomain=finance.yahoo.com&symbols="

//Get get quotes
func Get(s depot.Stock) (result Result) {
	// res := yahoo.Result{}

	u := fmt.Sprintf(url+"%v", s.Symbol)

	resp, err := http.Get(u)
	if err != nil {
		log.Printf("Error %v ", err)
		return result
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return result
	}

	// n := len(body)
	// txt := string(body[:n])

	// fmt.Println("Result for ", s.Name, "is ", txt)

	r, err := convertJSON2Result(body)
	if err != nil {
		log.Printf("Error during conversion from json to quote result. Stock: %v. Error: %v", s, err)
		return result
	}
	return r
}

func convertJSON2Result(b []byte) (result Result, err error) {
	out := QuoteResponse{}
	err = json.Unmarshal(b, &out)
	if err != nil {
		log.Println(err.Error())
		return result, fmt.Errorf("Error during json unmarshalling")
	}

	if len(out.QR.Res) == 0 {
		return result, fmt.Errorf("Can not find quotes")
	}

	return out.QR.Res[0], nil
}
