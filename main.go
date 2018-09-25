package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Nimsaja/DepotPerformance/depot"
)

var url = "https://query1.finance.yahoo.com/v7/finance/quote?lang=en-US&region=US&corsDomain=finance.yahoo.com&symbols="
var euro float32

func main() {
	start := time.Now()

	// euro, _ = getQuote(depot.Stock{Symbol: "USD-EUR", Count: 1})

	//to check for raceconditions -> go run -race main.go
	//to check if every go routine is done
	wg := sync.WaitGroup{}
	//need to tell the wait group how many go routines we have
	wg.Add(len(depot.Get()))
	for _, s := range depot.Get() {
		go func(s depot.Stock) {
			//will be called after this func is done, no matter where
			defer wg.Done()
			v := getQuote(s)
			fmt.Printf("Value for %v is %v\n", s.Name, v)
		}(s)
	}

	//here we wait for all the go routines to be done
	wg.Wait()
	fmt.Println("Elapsed Time ", time.Now().Sub(start))
}

func getQuote(s depot.Stock) (value float32) {
	u := fmt.Sprintf(url+"%v", s.Symbol)

	resp, err := http.Get(u)
	if err != nil {
		log.Printf("Error %v ", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return 0
	}

	n := len(body)
	txt := string(body[:n])

	fmt.Println("Result for ", s.Name, "is ", txt)

	// result := json.NewDecoder(body)

	// fmt.Println("Result for ", s.Name, "is ", result)

	return 107
}
