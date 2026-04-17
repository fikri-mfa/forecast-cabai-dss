package services

import (
	"errors"

	"forecast-cabai-dss/internal/domain"
	"forecast-cabai-dss/internal/forecasting"
	"forecast-cabai-dss/internal/repositories"
)

type ForecastService struct {
	hargaRepo repositories.HargaRepository
}

func NewForecastService(hargaRepo repositories.HargaRepository) *ForecastService {
	return &ForecastService{
		hargaRepo: hargaRepo,
	}
}

func (s *ForecastService) CalculateForecast(
	param domain.ParameterTES,
	forecastPeriods int,
) ([]float64, error) {

	data, err := s.hargaRepo.GetAllHarga()
	if err != nil {
		return nil, err
	}

	if len(data) < param.SeasonLength {
		return nil, errors.New("data historis tidak cukup untuk forecasting")
	}

	result := forecasting.TripleExponentialSmoothing(
		data,
		param.Alpha,
		param.Beta,
		param.Gamma,
		param.SeasonLength,
		forecastPeriods,
	)

	return result.Forecast, nil
}
