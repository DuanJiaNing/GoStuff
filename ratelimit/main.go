package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

// test script:
// for i in {1..6}; do curl http://localhost:8080/ping; done

var tlbthLimiter *limiter.Limiter

func init() {
	tlbthLimiter = tollbooth.NewLimiter(6, nil)
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage("rate limit reached\n")
}

func main() {
	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endpointHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Println(tollbooth.BuildKeys(tlbthLimiter, r))
	err := json.NewEncoder(w).Encode("hello world")
	if err != nil {
		return
	}
}
