package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Nimsaja/DepotPerformance/lima"

	"github.com/Nimsaja/DepotPerformance/store"

	"github.com/Nimsaja/DepotPerformance/depot"
	"github.com/Nimsaja/DepotPerformance/yahoo"
)

func main() {

	svl := lima.New(len(depot.Get()))
	start := time.Now()

	for _, s := range depot.Get() {
		go func(s depot.Stock) {
			//will be called after this func is done, no matter where
			defer svl.Done()

			v := yahoo.Get(s)
			svl.Add(lima.StockValue{Close: v.Close, Price: v.Price, Count: s.Count})
		}(s)
	}

	svl.Wait()

	fmt.Println("Elapsed Time ", time.Now().Sub(start))
	fmt.Println("Yesterday:", svl.SumYesterday(), "Today:", svl.SumToday(), "-> GAIN:", svl.Gain())

	// without store and paint graph
}

func _main() {
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
			quotesYesterday <- v.Close * s.Count
			quotesToday <- v.Price * s.Count
		}(s)
	}

	//here we wait for all the go routines to be done
	wg.Wait()
	fmt.Println("Elapsed Time ", time.Now().Sub(start))
	fmt.Println("Count:", len(depot.Get()))

	close(quotesYesterday)
	close(quotesToday)

	store.File(quotesYesterday)
	store.CreateGraph(quotesToday)
}
