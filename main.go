package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Nimsaja/DepotPerformance/depot"
	"github.com/Nimsaja/DepotPerformance/yahoo"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var url = "https://query1.finance.yahoo.com/v7/finance/quote?lang=en-US&region=US&corsDomain=finance.yahoo.com&symbols="
var euro float32
var quotes []float32
var date time.Time
var path = "stocksData.txt"

func main() {
	quotes = make([]float32, 0)

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
			quotes = append(quotes, v.Close*s.Count)
			date = time.Unix(v.Time, 0)
		}(s)
	}

	//here we wait for all the go routines to be done
	wg.Wait()
	fmt.Println("Elapsed Time ", time.Now().Sub(start))

	store()
	createGraph()
}

func createGraph() {
	//read in file
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	ax := make([]int, 0)
	ay := make([]float32, 0)
	var s []string
	var v float64
	nb := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := scanner.Text()
		fmt.Println(b)

		s = strings.Split(b, ", ")
		v, err = strconv.ParseFloat(s[1], 32)
		if err != nil {
			panic(err)
		}

		fmt.Println(v)

		t, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Println("Error parsing time ", err)
		}
		fmt.Println("Time: ", t)

		ax = append(ax, t)
		ay = append(ay, float32(v))

		nb++
	}

	//create plot
	xticks := plot.TimeTicks{Format: "2006-01-02"}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "My Depot Performance"
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Value"
	p.X.Tick.Marker = xticks

	pts := make(plotter.XYs, len(ay))
	for i, v := range ay {
		pts[i].X = float64(ax[i])
		pts[i].Y = float64(v)
	}

	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		panic(err)
	}

	//create x Axis
	fmt.Println(ax)

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "DepotPerformance.png"); err != nil {
		panic(err)
	}

	fmt.Println("\n*****Please open DepotPerformance.png********")
}

func store() {
	dummy := false
	var s string

	//create some dummy values if file does not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		dummy = true
	}

	//append to output file
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("Error %s ", err)
		panic(err)
	}

	defer f.Close()

	//TODO remove this later!!!!!!
	var t time.Time
	var n float64
	if dummy {
		for i := 30; i > 0; i-- {
			t = date.Add(time.Duration(-i) * time.Hour * 24)
			n = (rand.Float64() * 500) + 3000
			s = fmt.Sprintf("%v, %v", t.Unix(), n)
			fmt.Fprintln(f, s)
		}
	}

	s = fmt.Sprintf("%v, %v", date.Unix(), sum(quotes))
	fmt.Fprintln(f, s)
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
