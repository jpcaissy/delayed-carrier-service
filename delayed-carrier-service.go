package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Rate struct {
	ServiceName string `json:"service_name"`
	ServiceCode string `json:"service_code"`
	TotalPrice  int    `json:"total_price"`
	Currency    string `json:"currency"`
}

type Rates struct {
	Rates []Rate `json:"rates"`
}

func sleep_and_return_rates(w http.ResponseWriter, r *http.Request) {
	milliseconds := rand.Intn(7000)
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)

	rates := Rates{
		[]Rate{
			{
				"Standard rate",
				"standard-rate-1",
				rand.Intn(2000),
				"USD",
			},
			{
				"Expedited rate",
				"expedited-rate-2",
				rand.Intn(2000),
				"USD",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json_data, error := json.Marshal(rates)
	if error != nil {
		fmt.Println("Not enable to convert rates to json")
		io.WriteString(w, "{}")
		return
	}
	log.Println("Rendered shipping rates after " + strconv.Itoa(milliseconds) + " msec")

	w.Write(json_data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", sleep_and_return_rates)
	http.ListenAndServe(":"+port, nil)
}
