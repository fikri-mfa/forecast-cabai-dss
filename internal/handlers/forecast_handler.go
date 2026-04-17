package handlers

import (
	"encoding/json"
	"net/http"

	"forecast-cabai-dss/internal/domain"
	"forecast-cabai-dss/internal/services"
)

type ForecastHandler struct {
	forecastService *services.ForecastService
}

func NewForecastHandler(fs *services.ForecastService) *ForecastHandler {
	return &ForecastHandler{
		forecastService: fs,
	}
}

func (h *ForecastHandler) Forecast(w http.ResponseWriter, r *http.Request) {
	forecast, err := h.forecastService.CalculateForecast(
		domain.ParameterTES{
			Alpha:        0.2,
			Beta:         0.1,
			Gamma:        0.1,
			SeasonLength: 12,
		},
		3,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "success",
		"forecast": forecast,
	})
}
