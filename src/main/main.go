package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/duration", func(w http.ResponseWriter, r *http.Request) {
		result := map[string]string{
			"from": "Thursday, 4 January 2018",
			"to":   "Monday, 4 June 2018",
			"days": "152 days",
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
