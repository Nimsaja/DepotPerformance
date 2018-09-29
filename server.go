package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nimsaja/DepotPerformance/store"
	"github.com/Nimsaja/DepotPerformance/yahoo"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

// handle CORS and the OPTION method
func corsAndOptionHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			h.ServeHTTP(w, r)
		}
	}
}

// create all used Handler
func handler() http.Handler {
	router := mux.NewRouter()

	url := "/depot"
	router.HandleFunc(url, depot).Methods("GET")

	return corsAndOptionHandler(router)
}

func main() {

	http.Handle("/", handler())

	fmt.Println("*******Open http://localhost:8080/depot*******")
	fmt.Println()
	appengine.Main()
}

func depot(w http.ResponseWriter, r *http.Request) {
	//run go routines to get the stocks and calculate the sums
	yahoo.Run()

	//get the quotes for the client
	s := store.Get()

	writeOutAsJSON(w, s)
}

func writeOutAsJSON(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", string(b))
}
