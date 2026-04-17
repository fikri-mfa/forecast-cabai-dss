package services

import (
	"forecast-cabai-dss/internal/domain"
	"forecast-cabai-dss/internal/forecasting"
	"forecast-cabai-dss/internal/repositories"
)

type ForecastService struct {
	hargaRepo    *repositories.HargaPostgresRepository
	forecastRepo *repositories.ForecastPostgresRepository
}

func NewForecastService(
	hargaRepo *repositories.HargaPostgresRepository,
	forecastRepo *repositories.ForecastPostgresRepository,
) *ForecastService {
	return &ForecastService{
		hargaRepo:    hargaRepo,
		forecastRepo: forecastRepo,
	}
}

func (s *ForecastService) CalculateForecast(userID int, params domain.ParameterTES, periods int) (domain.ForecastResponse, error) {
	harga, err := s.hargaRepo.GetAllHarga()
	if err != nil {
		return domain.ForecastResponse{}, err
	}

	tesResult := forecasting.TripleExponentialSmoothing(
		harga,
		params.Alpha,
		params.Beta,
		params.Gamma,
		params.SeasonLength,
		periods,
	)

	mape := forecasting.MAPE(harga, tesResult.Fitted)
	rmse := forecasting.RMSE(harga, tesResult.Fitted)

	if err := s.forecastRepo.SaveForecast(userID, params, periods, tesResult.Forecast); err != nil {
		return domain.ForecastResponse{}, err
	}

	return domain.ForecastResponse{
		Forecast: tesResult.Forecast,
		Evaluation: domain.EvaluationResult{
			MAPE: mape,
			RMSE: rmse,
		},
	}, nil
}

func (s *ForecastService) GetForecastsByUserID(userID int) ([]domain.Forecast, error) {
	return s.forecastRepo.GetForecastsByUserID(userID)
}