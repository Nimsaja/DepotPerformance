package main

import (
	"fmt"
	"time"

	"github.com/Nimsaja/DepotPerformance/store"

	"github.com/Nimsaja/DepotPerformance/depot"
	"github.com/Nimsaja/DepotPerformance/yahoo"
)

func main() {
	//If some day I have a client which can add stocks, this is not necessary anymore
	depot.InitializeWithDefaultStocks()

	svl := depot.New(len(depot.Get()))
	start := time.Now()

	for _, s := range depot.Get() {
		go func(s depot.Stock) {
			defer svl.Done()

			v := yahoo.Get(s)
			svl.Add(depot.StockValue{Close: v.Close, Price: v.Price, Stock: s})
		}(s)
	}

	//here we wait for all the go routines to be done
	svl.Wait()

	fmt.Println("Elapsed Time ", time.Now().Sub(start))

	store.File(svl.SumYesterday())
	store.CreateGraph(svl.SumToday())
}
