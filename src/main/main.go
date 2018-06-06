package main

import (
	"dateCalculate"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/duration", func(w http.ResponseWriter, r *http.Request) {
		queryString := r.URL.Query()
		startDay := queryString.Get("start_day")
		startMonth := queryString.Get("start_month")
		startYear := queryString.Get("start_year")
		endDay := queryString.Get("end_day")
		endMonth := queryString.Get("end_month")
		endYear := queryString.Get("end_year")

		startDate := dateCalculate.StringToTime(startYear, startMonth, startDay)
		endDate := dateCalculate.StringToTime(endYear, endMonth, endDay)
		if startDate.After(endDate) {
			startDate, endDate = endDate, startDate
		}

		result := dateCalculate.MakeJson(startDate, endDate)
		resultJSON, err := json.Marshal(result)

		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(resultJSON)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
