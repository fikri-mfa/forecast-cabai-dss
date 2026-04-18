package routes

import (
    "forecast-cabai-dss/internal/handlers"
    "forecast-cabai-dss/internal/middlewares"
    "net/http"

    httpSwagger "github.com/swaggo/http-swagger"
    _ "forecast-cabai-dss/docs"
)

func RegisterRoutes(
    forecastHandler *handlers.ForecastHandler,
    authHandler *handlers.AuthHandler,
    hargaHandler *handlers.HargaHandler,
    dashboardHandler *handlers.DashboardHandler,
) http.Handler {
    mux := http.NewServeMux()

    // Swagger UI
    mux.Handle("/swagger/", httpSwagger.WrapHandler)

    // Public routes
    mux.HandleFunc("/register", authHandler.Register)
    mux.HandleFunc("/login", authHandler.Login)

    // Protected — Forecast
    mux.Handle("/forecast", middlewares.AuthMiddleware(http.HandlerFunc(forecastHandler.Forecast)))
    mux.Handle("/forecast/history", middlewares.AuthMiddleware(http.HandlerFunc(forecastHandler.GetHistory)))

    // Protected — Harga CRUD
    mux.Handle("/harga", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            hargaHandler.GetAllHarga(w, r)
        case http.MethodPost:
            hargaHandler.CreateHarga(w, r)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    })))

    mux.Handle("/harga/", middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPut:
            hargaHandler.UpdateHarga(w, r)
        case http.MethodDelete:
            hargaHandler.DeleteHarga(w, r)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    })))

    // Protected — Dashboard
    mux.Handle("/dashboard/stats", middlewares.AuthMiddleware(http.HandlerFunc(dashboardHandler.GetStats)))
    mux.Handle("/dashboard/recommendation", middlewares.AuthMiddleware(http.HandlerFunc(dashboardHandler.GetRecommendation)))

    // PENTING: chart/all harus didaftarkan SEBELUM chart
    mux.Handle("/dashboard/chart/all", middlewares.AuthMiddleware(http.HandlerFunc(dashboardHandler.GetAllChartData)))
    mux.Handle("/dashboard/chart", middlewares.AuthMiddleware(http.HandlerFunc(dashboardHandler.GetChartData)))
        return middlewares.CORSMiddleware(mux)
}