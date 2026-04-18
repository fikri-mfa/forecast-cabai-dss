package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"forecast-cabai-dss/internal/domain"
)

type ForecastPostgresRepository struct {
	db *sql.DB
}

func NewForecastPostgresRepository(db *sql.DB) *ForecastPostgresRepository {
	return &ForecastPostgresRepository{db: db}
}

func (r *ForecastPostgresRepository) SaveForecast(userID int, params domain.ParameterTES, periods int, result []float64) (int, error) {
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return 0, err
	}
	query := `
		INSERT INTO forecasts (user_id, alpha, beta, gamma, season_length, periods, result)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	var forecastID int
	err = r.db.QueryRow(query,
		userID, params.Alpha, params.Beta, params.Gamma,
		params.SeasonLength, periods, resultJSON,
	).Scan(&forecastID)
	return forecastID, err
}

func (r *ForecastPostgresRepository) SaveTesPerhitungan(data []domain.TesPerhitungan) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
		INSERT INTO tes_perhitungan (forecast_id, periode, tanggal, harga_asli, level, trend, seasonal, forecast)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, d := range data {
		_, err := stmt.Exec(
			d.ForecastID, d.Periode, d.Tanggal, d.HargaAsli,
			d.Level, d.Trend, d.Seasonal, d.Forecast,
		)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *ForecastPostgresRepository) GetForecastsByUserID(userID int) ([]domain.Forecast, error) {
	query := `
		SELECT id, user_id, alpha, beta, gamma, season_length, periods, result, created_at
		FROM forecasts
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forecasts []domain.Forecast
	for rows.Next() {
		var f domain.Forecast
		var resultJSON []byte
		err := rows.Scan(
			&f.ID, &f.UserID, &f.Alpha, &f.Beta, &f.Gamma,
			&f.SeasonLength, &f.Periods, &resultJSON, &f.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(resultJSON, &f.Result); err != nil {
			return nil, err
		}
		forecasts = append(forecasts, f)
	}
	return forecasts, nil
}

func (r *ForecastPostgresRepository) CountForecasts() (int, error) {
	var count int
	err := r.db.QueryRowContext(context.Background(),
		`SELECT COUNT(*) FROM forecasts`).Scan(&count)
	return count, err
}

// MAPE diambil dari tabel tes_perhitungan karena result hanya simpan []float64
func (r *ForecastPostgresRepository) GetAverageMAPE(userID int) (float64, error) {
	var avg float64
	err := r.db.QueryRowContext(context.Background(), `
		SELECT COALESCE(AVG(
			ABS(tp.harga_asli - tp.forecast) / NULLIF(tp.harga_asli, 0) * 100
		), 0)
		FROM tes_perhitungan tp
		JOIN forecasts f ON f.id = tp.forecast_id
		WHERE f.user_id = $1
		AND tp.forecast IS NOT NULL
		AND tp.harga_asli > 0
	`, userID).Scan(&avg)
	return avg, err
}