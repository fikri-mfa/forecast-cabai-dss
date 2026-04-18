package handlers

import (
	"encoding/json"
	"forecast-cabai-dss/internal/auth"
	"forecast-cabai-dss/internal/domain"
	"forecast-cabai-dss/internal/middlewares"
	"forecast-cabai-dss/internal/services"
	"net/http"
)

type ForecastHandler struct {
	forecastService *services.ForecastService
}

func NewForecastHandler(fs *services.ForecastService) *ForecastHandler {
	return &ForecastHandler{forecastService: fs}
}

// Forecast godoc
// @Summary      Hitung forecast harga cabai
// @Description  Menghitung prediksi harga cabai menggunakan Triple Exponential Smoothing
// @Tags         Forecast
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body domain.ForecastRequest false "Parameter TES (opsional, ada default value)"
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /forecast [post]
func (h *ForecastHandler) Forecast(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "tidak bisa ambil data user", http.StatusUnauthorized)
		return
	}

	var req domain.ForecastRequest
	autoOptimize := false

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Alpha == 0 {
		req = domain.ForecastRequest{
			SeasonLength: 12,
			Periods:      3,
		}
		autoOptimize = true
	}

	response, err := h.forecastService.CalculateForecast(
		claims.UserID,
		domain.ParameterTES{
			Alpha:        req.Alpha,
			Beta:         req.Beta,
			Gamma:        req.Gamma,
			SeasonLength: req.SeasonLength,
		},
		req.Periods,
		autoOptimize,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":         "success",
		"forecast":       response.Forecast,
		"auto_optimized": response.AutoOptimized,
		"params_used":    response.ParamsUsed,
		"evaluation": map[string]interface{}{
			"mape": response.Evaluation.MAPE,
			"rmse": response.Evaluation.RMSE,
		},
		"perhitungan": response.Perhitungan,
	})
}

// GetHistory godoc
// @Summary      Ambil riwayat forecast
// @Description  Mengambil semua riwayat forecast milik user yang sedang login
// @Tags         Forecast
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /forecast/history [get]
func (h *ForecastHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserClaimsKey).(*auth.Claims)
	if !ok {
		http.Error(w, "tidak bisa ambil data user", http.StatusUnauthorized)
		return
	}

	forecasts, err := h.forecastService.GetForecastsByUserID(claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   forecasts,
	})
}