package handlers

import (
	"encoding/json"
	"forecast-cabai-dss/internal/auth"
	"forecast-cabai-dss/internal/middlewares"
	"forecast-cabai-dss/internal/services"
	"net/http"
)

type DashboardHandler struct {
	dashboardService *services.DashboardService
}

func NewDashboardHandler(ds *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{dashboardService: ds}
}

// GET /dashboard/stats
func (h *DashboardHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	stats, err := h.dashboardService.GetStats(claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   stats,
	})
}

// GET /dashboard/chart
func (h *DashboardHandler) GetChartData(w http.ResponseWriter, r *http.Request) {
	data, err := h.dashboardService.GetChartData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

// GET /dashboard/recommendation
func (h *DashboardHandler) GetRecommendation(w http.ResponseWriter, r *http.Request) {
	harga, err := h.dashboardService.GetRecommendation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   map[string]float64{"rekomendasi": harga},
	})
}

// GET /dashboard/chart/all
func (h *DashboardHandler) GetAllChartData(w http.ResponseWriter, r *http.Request) {
	data, err := h.dashboardService.GetAllChartData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}