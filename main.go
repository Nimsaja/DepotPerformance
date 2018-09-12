package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Nimsaja/DepotPerformance/depot"
)

var url = "http://markets.financialcontent.com/stocks/action/gethistoricaldata?Symbol="
var euro float32

func main() {
	euro = getClose(depot.Stock{Symbol: "USD-EUR", Count: 1})

	var v float32
	for _, s := range depot.Get() {
		v = getClose(s)
		fmt.Printf("Value for %v in Euro is %v \n", s.Name, v*euro)
	}
}

func getClose(s depot.Stock) float32 {
	resp, err := http.Get(url + s.Symbol)
	if err != nil {
		log.Printf("Error %v ", err)
		return 0
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	//skip first line
	line, err := reader.Read()
	line, err = reader.Read()

	f, err := strconv.ParseFloat(line[5], 32)
	if err != nil {
		log.Printf("Error %v ", err)
		return 0
	}

	return float32(f) * s.Count
}
