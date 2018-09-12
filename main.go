package main

import (
	"encoding/csv"
	"log"
	"net/http"
)

func main() {

	urlGoogle := "http://markets.financialcontent.com/stocks/action/gethistoricaldata?Symbol=GOOG"
	// urlAmazon := "http://markets.financialcontent.com/stocks/action/gethistoricaldata?Symbol=AMAZ"

	resp, err := http.Get(urlGoogle)
	if err != nil {
		log.Printf("Error %v ", err)
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	//skip first line
	line, err := reader.Read()
	line, err = reader.Read()

	log.Printf("Close from yesterday %v ", line[5])
}
