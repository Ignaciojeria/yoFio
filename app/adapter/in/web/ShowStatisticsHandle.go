package web

import (
	"encoding/json"
	"net/http"
	"yofio/app/business/port/out"
)

type ShowStatisticsHandle struct {
	out.LoadStatisticsPortOUT
}

func (h ShowStatisticsHandle) ShowStatistics(w http.ResponseWriter, req *http.Request) {
	statistics := h.LoadStatisticsPortOUT()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statistics)
}
