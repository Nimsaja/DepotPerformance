package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Nimsaja/DepotPerformance/store"

	"github.com/Nimsaja/DepotPerformance/depot"
	"github.com/Nimsaja/DepotPerformance/yahoo"
)

func main() {
	//If some day I have a client which can add stocks, this is not necessary anymore
	depot.InitializeWithDefaultStocks()

	//declaration of channel
	quotesYesterday := make(chan float32, len(depot.Get()))
	quotesToday := make(chan float32, len(depot.Get()))

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

			v := yahoo.Get(s)

			c := v.Close * s.Count
			p := v.Price * s.Count
			quotesYesterday <- c
			quotesToday <- p

			// fmt.Println(s.Name, ": ", c, " -> ", p)
			// fmt.Println(s.Name, ": ", c-s.Buy*s.Count, " -> ", p-s.Buy*s.Count)
		}(s)
	}

	//here we wait for all the go routines to be done
	wg.Wait()
	fmt.Println("Elapsed Time ", time.Now().Sub(start))

	close(quotesYesterday)
	close(quotesToday)

	store.File(quotesYesterday)
	store.CreateGraph(quotesToday)
}
