package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Nimsaja/DepotPerformance/depot"
)

var url = "http://markets.financialcontent.com/stocks/action/gethistoricaldata?"

func main() {
	start := time.Now()

	ch4Euro := make(chan float32, 0)

	go func() {
		euro, _ := getClose(depot.Stock{Symbol: "USD-EUR", Count: 1})
		ch4Euro <- euro

		close(ch4Euro)
	}()

	//need to create a new struct
	type closureValues struct {
		s depot.Stock
		d string
		v float32
	}

	lenDepot := len(depot.Get())

	//declaration of channel
	ch4Stocks := make(chan closureValues, lenDepot)

	//to check for raceconditions -> go run -race main.go
	//to check if every go routine is done
	wg := sync.WaitGroup{}

	//need to tell the wait group how many go routines we have
	wg.Add(lenDepot)
	for _, s := range depot.Get() {
		go func(s depot.Stock) {
			//will be called after this func is done, no matter where
			defer wg.Done()
			//input to channel
			v, d := getClose(s)
			ch4Stocks <- closureValues{s: s, d: d, v: v}
		}(s)
	}
	//here we wait for all the go routines to be done
	wg.Wait()

	close(ch4Stocks)

	euro := <-ch4Euro

	//output of channel
	for cV := range ch4Stocks {
		fmt.Printf("Value for %v on %v is %v Euro\n", cV.s.Name, cV.d, cV.s.AsEuro(cV.v, euro))
	}

	fmt.Println("Elapsed Time ", time.Now().Sub(start))
}

func getClose(s depot.Stock) (value float32, date string) {
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
