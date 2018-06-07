package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func durationHandler(w http.ResponseWriter, r *http.Request) {
	durationResponse := map[string]string{
		"from":        "Thursday, 4 January 2018",
		"to":          "Monday, 4 June 2018",
		"days":        "152 days",
		"years":       "5 months, 1 day",
		"seconds":     "13,132,800 seconds",
		"minutes":     "218,880 minutes",
		"hours":       "3,648 hours",
		"weeks":       "21 weeks and 5 days",
		"ratioOfYear": "41.64% of 2018",
	}

	err := json.NewEncoder(w).Encode(durationResponse)
	if err != nil {
		panic(err)
	}

}
func main() {
	http.HandleFunc("/duration", durationHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
