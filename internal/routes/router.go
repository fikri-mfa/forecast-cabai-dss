package routes

import (
	"forecast-cabai-dss/internal/handlers"
	"forecast-cabai-dss/internal/middlewares"
	"net/http"
)

func RegisterRoutes(
	forecastHandler *handlers.ForecastHandler,
	authHandler *handlers.AuthHandler,
) http.Handler {
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/register", authHandler.Register)
	mux.HandleFunc("/login", authHandler.Login)

	// Protected routes
	mux.Handle("/forecast", middlewares.AuthMiddleware(http.HandlerFunc(forecastHandler.Forecast)))

	return mux
}