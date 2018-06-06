package main

import (
	"dateCalculate"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		duration := dateCalculate.Diff(startDate, endDate)

		result := map[string]string{
			"from": dateCalculate.FormatDate(startDate),
			"to":   dateCalculate.FormatDate(endDate),
			"days": strconv.Itoa(duration) + " days",
		}
		resultJSON, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(resultJSON)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
