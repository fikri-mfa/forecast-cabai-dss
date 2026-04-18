package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"forecast-cabai-dss/internal/config"
	"forecast-cabai-dss/internal/handlers"
	"forecast-cabai-dss/internal/repositories"
	"forecast-cabai-dss/internal/routes"
	"forecast-cabai-dss/internal/services"
)

func main() {
	godotenv.Load()

	db, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	// Repositories
	hargaRepo := repositories.NewHargaPostgresRepository(db)
	userRepo := repositories.NewUserPostgresRepository(db)
	forecastRepo := repositories.NewForecastPostgresRepository(db)

	// Services
	forecastService := services.NewForecastService(hargaRepo, forecastRepo)
	authService := services.NewAuthService(userRepo)
	hargaService := services.NewHargaService(hargaRepo)
	dashboardService := services.NewDashboardService(hargaRepo, forecastRepo)

	// Handlers
	forecastHandler := handlers.NewForecastHandler(forecastService)
	authHandler := handlers.NewAuthHandler(authService)
	hargaHandler := handlers.NewHargaHandler(hargaService)
	dashboardHandler := handlers.NewDashboardHandler(dashboardService)

	router := routes.RegisterRoutes(forecastHandler, authHandler, hargaHandler, dashboardHandler)

	log.Println("Server running at :9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}