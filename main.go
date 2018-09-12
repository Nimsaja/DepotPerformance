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

func main() {
	var v float64
	for _, s := range depot.Get() {
		v = getClose(s.Symbol)
		fmt.Printf("Price for %v in Dollar is %v \n", s.Name, v)

		v = convertToEuro(v)
		fmt.Printf("Price for %v in Euro is %v \n", s.Name, v)

	}
}

func getClose(s string) float64 {
	resp, err := http.Get(url + s)
	if err != nil {
		log.Printf("Error %v ", err)
		return 0
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	//skip first line
	line, err := reader.Read()
	line, err = reader.Read()

	f, err := strconv.ParseFloat(line[5], 64)
	if err != nil {
		log.Printf("Error %v ", err)
		return 0
	}

	return f
}

func convertToEuro(v float64) float64 {
	e := getClose("USD-EUR")

	return v * e
}
