package handlers

import (
	"net/http"
	"personalwebsite/utils"
)

func HandleMetrics(w http.ResponseWriter, r *http.Request) {
	metricsJSON, err := utils.GetMetricsJSON()
	if err != nil {
		http.Error(w, "Error calculating metrics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(metricsJSON)
}
