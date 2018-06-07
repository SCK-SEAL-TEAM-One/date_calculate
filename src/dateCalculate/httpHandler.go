package dateCalculate

import (
	"encoding/json"
	"net/http"
)

func DurationHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	startDay := queryString.Get("start_day")
	startMonth := queryString.Get("start_month")
	startYear := queryString.Get("start_year")
	endDay := queryString.Get("end_day")
	endMonth := queryString.Get("end_month")
	endYear := queryString.Get("end_year")

	startDate := StringToDate(startDay, startMonth, startYear)
	endDate := StringToDate(endDay, endMonth, endYear)

	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}

	err := json.NewEncoder(w).Encode(MakeJSON(startDate, endDate))
	if err != nil {
		panic(err)
	}
}
