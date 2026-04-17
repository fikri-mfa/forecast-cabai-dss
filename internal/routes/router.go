package routes

import (
    "forecast-cabai-dss/internal/handlers"
    "forecast-cabai-dss/internal/middlewares"
    "net/http"

    httpSwagger "github.com/swaggo/http-swagger"
    _ "forecast-cabai-dss/docs" // hasil generate swag init
)

func RegisterRoutes(
    forecastHandler *handlers.ForecastHandler,
    authHandler *handlers.AuthHandler,
) http.Handler {
    mux := http.NewServeMux()

    // Swagger UI
    mux.Handle("/swagger/", httpSwagger.WrapHandler)

    // Public routes
    mux.HandleFunc("/register", authHandler.Register)
    mux.HandleFunc("/login", authHandler.Login)

    // Protected routes
    mux.Handle("/forecast", middlewares.AuthMiddleware(http.HandlerFunc(forecastHandler.Forecast)))
    mux.Handle("/forecast/history", middlewares.AuthMiddleware(http.HandlerFunc(forecastHandler.GetHistory)))

    return middlewares.CORSMiddleware(mux)
}