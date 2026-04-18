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

func (s *ForecastService) CalculateForecast(userID int, params domain.ParameterTES, periods int, autoOptimize bool) (domain.ForecastResponse, error) {
	// 1. Ambil data harga
	hargaData, err := s.hargaRepo.GetAllHarga()
	if err != nil {
		return domain.ForecastResponse{}, err
	}

	// 2. Cari parameter optimal kalau autoOptimize = true
	autoOptimized := false
	if autoOptimize {
		optimal := forecasting.FindOptimalParams(hargaData, params.SeasonLength)
		params.Alpha = optimal.Alpha
		params.Beta = optimal.Beta
		params.Gamma = optimal.Gamma
		autoOptimized = true
	}

	// 3. Jalankan TES
	tesResult := forecasting.TripleExponentialSmoothing(
		hargaData,
		params.Alpha,
		params.Beta,
		params.Gamma,
		params.SeasonLength,
		periods,
	)

	// 4. Hitung evaluasi (hanya dari data di luar periode inisialisasi)
	mape := forecasting.MAPE(hargaData[params.SeasonLength:], tesResult.Fitted[params.SeasonLength:])
	rmse := forecasting.RMSE(hargaData[params.SeasonLength:], tesResult.Fitted[params.SeasonLength:])

	// 5. Simpan ke tabel forecasts
	forecastID, err := s.forecastRepo.SaveForecast(userID, params, periods, tesResult.Forecast)
	if err != nil {
		return domain.ForecastResponse{}, err
	}

	// 6. Bangun detail perhitungan per periode
	hargaRows, err := s.hargaRepo.GetAllHargaWithDate()
	if err != nil {
		return domain.ForecastResponse{}, err
	}

	perhitungan := make([]domain.TesPerhitungan, len(hargaRows))
	for i, row := range hargaRows {
		p := domain.TesPerhitungan{
			ForecastID: forecastID,
			Periode:    i + 1,
			Tanggal:    row.Tanggal,
			HargaAsli:  row.Harga,
		}

		if i < params.SeasonLength {
			// Periode awal — level, trend, seasonal ada tapi forecast NULL
			l := tesResult.Level[i]
			t := tesResult.Trend[i]
			ss := tesResult.Seasonal[i]
			p.Level = &l
			p.Trend = &t
			p.Seasonal = &ss
			p.Forecast = nil
		} else {
			l := tesResult.Level[i]
			t := tesResult.Trend[i]
			ss := tesResult.Seasonal[i]
			f := tesResult.Fitted[i]
			p.Level = &l
			p.Trend = &t
			p.Seasonal = &ss
			p.Forecast = &f
		}

		perhitungan[i] = p
	}

	// 7. Simpan ke tabel tes_perhitungan
	if err := s.forecastRepo.SaveTesPerhitungan(perhitungan); err != nil {
		return domain.ForecastResponse{}, err
	}

	return domain.ForecastResponse{
		Forecast:      tesResult.Forecast,
		Evaluation:    domain.EvaluationResult{MAPE: mape, RMSE: rmse},
		Perhitungan:   perhitungan,
		ParamsUsed:    params,
		AutoOptimized: autoOptimized,
	}, nil
}

func (s *ForecastService) GetForecastsByUserID(userID int) ([]domain.Forecast, error) {
	return s.forecastRepo.GetForecastsByUserID(userID)
}