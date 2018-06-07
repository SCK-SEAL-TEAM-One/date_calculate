package main

import (
	"dateCalculate"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/duration", dateCalculate.DurationHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
