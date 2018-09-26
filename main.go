package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Nimsaja/DepotPerformance/depot"
	"github.com/Nimsaja/DepotPerformance/yahoo"
)

var url = "https://query1.finance.yahoo.com/v7/finance/quote?lang=en-US&region=US&corsDomain=finance.yahoo.com&symbols="
var histURL = "https://query1.finance.yahoo.com/v7/finance/spark?symbols="
var histURLArgs = "&range=1mo&interval=1d"
var euro float32

func main() {
	quotesYesterday := make([]float32, 0)
	quotesToday := make([]float32, 0)
	quotesHist := make([]yahoo.Data, 0)

	start := time.Now()

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

			// fmt.Printf("Result for %v is %v\n", s.Name, v)

			//Want to sum here - race condition?!?
			quotesYesterday = append(quotesYesterday, v.Close*s.Count)
			quotesToday = append(quotesToday, v.Price*s.Count)

			vh := getHistQuote(s)

			q := yahoo.Data{S: s, TV: yahoo.CreateTimeToValueMap(vh.Resp[0].T, vh.Resp[0].I.Q[0].V)}
			quotesHist = append(quotesHist, q)
		}(s)
	}

	//here we wait for all the go routines to be done
	wg.Wait()
	fmt.Println("Elapsed Time ", time.Now().Sub(start))

	//actual data
	fmt.Println("Depot yesterday ", sum(quotesYesterday), " / Depot today ", sum(quotesToday))
	fmt.Println("-> Diff ", sum(quotesToday)-sum(quotesYesterday))

	//historical data
	fmt.Println("\n***********HistoricalData****************")
	for _, q := range quotesHist {
		fmt.Println("\n", q.S.Name, ": ")
		for k, v := range q.TV {
			fmt.Println(k, " -> ", v)
		}
	}

}

func sum(input []float32) float32 {
	sum := float32(0.0)

	for i := range input {
		sum += input[i]
	}

	return sum
}

func getQuote(s depot.Stock) (result yahoo.Result) {
	res := yahoo.Result{}

	u := fmt.Sprintf(url+"%v", s.Symbol)

	resp, err := http.Get(u)
	if err != nil {
		log.Printf("Error %v ", err)
		return res
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return res
	}

	// n := len(body)
	// txt := string(body[:n])

	// fmt.Println("Result for ", s.Name, "is ", txt)

	out := yahoo.QuoteResponse{}
	err = json.Unmarshal(body, &out)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(out.QR.Res) > 0 {
		res = out.QR.Res[0]
	}

	return res
}

func getHistQuote(s depot.Stock) (result yahoo.HistResult) {
	res := yahoo.HistResult{}

	u := fmt.Sprintf(histURL+"%v"+histURLArgs, s.Symbol)

	resp, err := http.Get(u)
	if err != nil {
		log.Printf("Error %v ", err)
		return res
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return res
	}

	// n := len(body)
	// txt := string(body[:n])
	// fmt.Println("Result for ", s.Name, "is ", txt)

	out := yahoo.Spark{}
	err = json.Unmarshal(body, &out)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if len(out.SP.Res) > 0 {
		res = out.SP.Res[0]
	}

	return res
}
