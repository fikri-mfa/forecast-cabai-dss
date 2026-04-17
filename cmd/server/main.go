// @title           Forecast Cabai DSS API
// @version         1.0
// @description     API untuk sistem pendukung keputusan forecast harga cabai menggunakan Triple Exponential Smoothing
// @host            localhost:9090
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token dengan format: Bearer {token}
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

	// Handlers
	forecastHandler := handlers.NewForecastHandler(forecastService)
	authHandler := handlers.NewAuthHandler(authService)

	router := routes.RegisterRoutes(forecastHandler, authHandler)

	log.Println("Server running at :9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}