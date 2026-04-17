package repositories

import (
	"context"
	"database/sql"
)

type HargaPostgresRepository struct {
	db *sql.DB
}

func NewHargaPostgresRepository(db *sql.DB) *HargaPostgresRepository {
	return &HargaPostgresRepository{
		db: db,
	}
}

func (r *HargaPostgresRepository) GetAllHarga() ([]float64, error) {
	query := `
		SELECT harga
		FROM harga
		ORDER BY tanggal ASC
	`

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prices []float64

	for rows.Next() {
		var price float64
		if err := rows.Scan(&price); err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return prices, nil
}
