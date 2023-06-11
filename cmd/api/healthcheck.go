package api

import (
	"encoding/json"
	"net/http"
)

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "OK",
	}

	jsonRes, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
