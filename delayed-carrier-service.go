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
	milliseconds := sleep()

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

	json_data, err := json.Marshal(rates)
	if err != nil {
		fmt.Println("unable to marshal json properly" + err.Error())
		io.WriteString(w, "{}")
		return
	}
	logSuccess("Rendered shipping rates after ", milliseconds)

	w.Write(json_data)
}

func delayedAuthy(w http.ResponseWriter, r *http.Request) {
	milliseconds := sleep()
	logSuccess("Rendered auty/verification/* successful response after ", milliseconds)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{}`))
}

func delayedAuthyPhoneIntelligence(w http.ResponseWriter, r *http.Request) {
	milliseconds := sleep()
	logSuccess("Rendered auty/info successful response after ", milliseconds)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success":true, "type": "cellphone", "ported": false}`))
}

func sleep() int {
	milliseconds := rand.Intn(7000)
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)

	return milliseconds
}

func logSuccess(message string, milliseconds int) {
	log.Println(message + strconv.Itoa(milliseconds) + " msec")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	http.HandleFunc("/", sleep_and_return_rates)
	http.HandleFunc("/authy/verification/start", delayedAuthy)
	http.HandleFunc("/authy/verification/check", delayedAuthy)
	http.HandleFunc("/authy/info", delayedAuthyPhoneIntelligence)
	http.ListenAndServe(":"+port, nil)
}
