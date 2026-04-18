package services

import (
	"forecast-cabai-dss/internal/domain"
	"forecast-cabai-dss/internal/forecasting"
	"forecast-cabai-dss/internal/repositories"
)

type DashboardService struct {
	hargaRepo    *repositories.HargaPostgresRepository
	forecastRepo *repositories.ForecastPostgresRepository
}

func NewDashboardService(
	hargaRepo *repositories.HargaPostgresRepository,
	forecastRepo *repositories.ForecastPostgresRepository,
) *DashboardService {
	return &DashboardService{hargaRepo: hargaRepo, forecastRepo: forecastRepo}
}

func (s *DashboardService) GetStats(userID int) (domain.DashboardStats, error) {
	totalHarga, err := s.hargaRepo.CountHarga()
	if err != nil {
		return domain.DashboardStats{}, err
	}

	totalForecast, err := s.forecastRepo.CountForecasts()
	if err != nil {
		return domain.DashboardStats{}, err
	}

	mape, err := s.forecastRepo.GetAverageMAPE(userID)
	if err != nil {
		return domain.DashboardStats{}, err
	}
	akurasi := 100 - mape

	rekomendasi, err := s.GetRecommendation()
	if err != nil {
		rekomendasi = 0
	}

	return domain.DashboardStats{
		TotalDataHarga:  totalHarga,
		TotalForecasts:  totalForecast,
		AkurasiSistem:   akurasi,
		RekomendasiHari: rekomendasi,
	}, nil
}

func (s *DashboardService) GetRecommendation() (float64, error) {
	data, err := s.hargaRepo.GetAllHarga()
	if err != nil {
		return 0, err
	}
	if len(data) < 12 {
		return 0, nil
	}

	optimal := forecasting.FindOptimalParams(data, 12)
	result := forecasting.TripleExponentialSmoothing(
		data, optimal.Alpha, optimal.Beta, optimal.Gamma, 12, 1,
	)
	if len(result.Forecast) == 0 {
		return 0, nil
	}
	return result.Forecast[0], nil
}

func (s *DashboardService) GetChartData() ([]domain.ChartData, error) {
	rows, err := s.hargaRepo.GetAllHargaWithDate()
	if err != nil {
		return nil, err
	}

	data := make([]float64, len(rows))
	for i, r := range rows {
		data[i] = r.Harga
	}

	if len(data) < 12 {
		return nil, nil
	}

	optimal := forecasting.FindOptimalParams(data, 12)
	result := forecasting.TripleExponentialSmoothing(
		data, optimal.Alpha, optimal.Beta, optimal.Gamma, 12, 0,
	)

	// Ambil 7 data terakhir
	n := len(rows)
	start := n - 7
	if start < 0 {
		start = 0
	}

	var chartData []domain.ChartData
	for i := start; i < n; i++ {
		fitted := 0.0
		if i < len(result.Fitted) {
			fitted = result.Fitted[i]
		}
		chartData = append(chartData, domain.ChartData{
			Tanggal:  rows[i].Tanggal,
			Aktual:   rows[i].Harga,
			Prediksi: fitted,
		})
	}
	return chartData, nil
}
func (s *DashboardService) GetAllChartData() ([]domain.ChartData, error) {
	rows, err := s.hargaRepo.GetAllHargaWithDate()
	if err != nil {
		return nil, err
	}

	data := make([]float64, len(rows))
	for i, r := range rows {
		data[i] = r.Harga
	}

	if len(data) < 12 {
		return nil, nil
	}

	optimal := forecasting.FindOptimalParams(data, 12)
	result := forecasting.TripleExponentialSmoothing(
		data, optimal.Alpha, optimal.Beta, optimal.Gamma, 12, 0,
	)

	var chartData []domain.ChartData
	for i, row := range rows {
		fitted := 0.0
		if i < len(result.Fitted) {
			fitted = result.Fitted[i]
		}
		chartData = append(chartData, domain.ChartData{
			Tanggal:  row.Tanggal,
			Aktual:   row.Harga,
			Prediksi: fitted,
		})
	}
	return chartData, nil
}