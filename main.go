package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Nimsaja/DepotPerformance/depot"
)

var url = "http://markets.financialcontent.com/stocks/action/gethistoricaldata?"
var euro float32

func main() {
	euro, _ = getClose(depot.Stock{Symbol: "USD-EUR", Count: 1})

	var v float32
	var d string
	for _, s := range depot.Get() {
		v, d = getClose(s)
		fmt.Printf("Value for %v on %v is %v Euro\n", s.Name, d, v*euro)
	}
}

func getClose(s depot.Stock) (float32, string) {
	// m := 8 //start month
	sy := s.Symbol
	r := 1 //how many month
	y := 2018

	// u := fmt.Sprintf(url+"Month=%v&Symbol=%v&Range=%v&Year=%v", m, sy, r, y)
	u := fmt.Sprintf(url+"Symbol=%v&Range=%v&Year=%v", sy, r, y)

	resp, err := http.Get(u)
	if err != nil {
		log.Printf("Error %v ", err)
		return 0, ""
	}
	defer resp.Body.Close()

	// var lineCount = 0
	reader := csv.NewReader(resp.Body)

	// for {
	// 	record, err := reader.Read()
	// 	// end-of-file is fitted into err
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return 0
	// 	}
	// 	fmt.Println(lineCount, " -> ", record)
	// 	lineCount++
	// }

	//skip first line
	line, err := reader.Read()
	line, err = reader.Read()

	f, err := strconv.ParseFloat(line[5], 32)
	if err != nil {
		log.Printf("Error %v ", err)
		return 0, ""
	}

	return float32(f) * s.Count, line[1]
}
