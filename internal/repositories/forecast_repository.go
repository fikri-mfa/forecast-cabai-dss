package repositories

import (
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

func (r *ForecastPostgresRepository) SaveForecast(userID int, params domain.ParameterTES, periods int, result []float64) error {
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO forecasts (user_id, alpha, beta, gamma, season_length, periods, result)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = r.db.Exec(query,
		userID,
		params.Alpha,
		params.Beta,
		params.Gamma,
		params.SeasonLength,
		periods,
		resultJSON,
	)
	return err
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