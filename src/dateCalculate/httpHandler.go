package dateCalculate

import (
	"encoding/json"
	"net/http"
	"time"
)

func DurationHandler(w http.ResponseWriter, r *http.Request) {
	startDate := time.Now()
	endDate := time.Now()

	err := json.NewEncoder(w).Encode(MakeJSON(startDate,endDate))
	if err != nil {
		panic(err)
	}
}
