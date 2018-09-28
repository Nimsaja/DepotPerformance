package store

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Nimsaja/DepotPerformance/depot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const path = "stocksData.txt"

//File store channel inputs to file
func File(ch chan float32) {
	//append to output file
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("Error %s ", err)
		panic(err)
	}

	defer f.Close()

	//output of channel - sum up
	var sum float32
	for v := range ch {
		sum += v
	}

	//date should be the close time from yesterday - say 23:59
	d := time.Now().Add(time.Duration(-1) * time.Hour * 24)
	d = time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 0, 0, time.UTC)
	s := fmt.Sprintf("%v, %v", d.Unix(), sum)
	fmt.Fprintln(f, s)

	f.Close()
}

//CreateGraph create Graph to show the depot value over time
func CreateGraph(ch chan float32) {
	//read in file
	f, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	at := make([]int, 0)
	av := make([]float32, 0)
	var s []string
	var v float64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		b := scanner.Text()
		s = strings.Split(b, ", ")

		//get time
		t, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Println("Error parsing time ", err)
		}

		at = append(at, t)

		//get value
		v, err = strconv.ParseFloat(s[1], 32)
		if err != nil {
			panic(err)
		}

		av = append(av, float32(v))
	}

	//remove duplicated times in file - can be removed as soon as storing values is automatically called once a day
	ax := make([]int, 0)
	ay := make([]float32, 0)
	prevTimes := make(map[int]struct{})
	var txt string
	for i, t := range at {
		_, exists := prevTimes[t]
		if !exists {
			ax = append(ax, t)
			ay = append(ay, av[i])

			txt = fmt.Sprintf("%v, %v", t, v)
			fmt.Fprintln(f, txt)

			prevTimes[t] = struct{}{}
		}
	}

	fmt.Println("Arrays ", ax, " / ", ay)

	//todays values
	var sum float32
	for value := range ch {
		sum += value
	}
	ax = append(ax, int(time.Now().Unix()))
	ay = append(ay, float32(sum))

	//create plots
	xticks := plot.TimeTicks{Format: "2006-01-02\n15:04"}

	//amount of depot
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

	// Save the plot to a PNG file.
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "DepotPerformance.png"); err != nil {
		panic(err)
	}

	fmt.Println("\n*****Please open DepotPerformance.png********")

	//diff of depot
	pd, err := plot.New()
	if err != nil {
		panic(err)
	}

	pd.Title.Text = "Depot Diff"
	pd.X.Label.Text = "Date"
	pd.Y.Label.Text = "Value"
	pd.X.Tick.Marker = xticks

	ptsd := make(plotter.XYs, len(ay))
	for i, v := range ay {
		ptsd[i].X = float64(ax[i])
		ptsd[i].Y = float64(v - depot.SumBuy())

		fmt.Printf("Plot: t=%v, v=%v, diff=%v\n", time.Unix(int64(ax[i]), 0), v, v-depot.SumBuy())
	}

	err = plotutil.AddLinePoints(pd, ptsd)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := pd.Save(10*vg.Inch, 4*vg.Inch, "DepotDiff.png"); err != nil {
		panic(err)
	}

	fmt.Println("\n*****Please open DepotDiff.png********")

	f.Close()
}
