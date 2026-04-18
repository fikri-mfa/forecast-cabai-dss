package repositories

import (
	"context"
	"database/sql"
	"forecast-cabai-dss/internal/domain"
)

type HargaPostgresRepository struct {
	db *sql.DB
}

func NewHargaPostgresRepository(db *sql.DB) *HargaPostgresRepository {
	return &HargaPostgresRepository{db: db}
}

func (r *HargaPostgresRepository) GetAllHarga() ([]float64, error) {
	rows, err := r.db.QueryContext(context.Background(),
		`SELECT harga FROM harga ORDER BY tanggal ASC`)
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
	return prices, rows.Err()
}

func (r *HargaPostgresRepository) GetAllHargaWithDate() ([]domain.HargaRow, error) {
	rows, err := r.db.QueryContext(context.Background(),
		`SELECT id, tanggal, harga FROM harga ORDER BY tanggal ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.HargaRow
	for rows.Next() {
		var row domain.HargaRow
		if err := rows.Scan(&row.ID, &row.Tanggal, &row.Harga); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, rows.Err()
}

func (r *HargaPostgresRepository) CreateHarga(tanggal string, harga float64) error {
	_, err := r.db.ExecContext(context.Background(),
		`INSERT INTO harga (tanggal, harga) VALUES ($1, $2)`, tanggal, harga)
	return err
}

func (r *HargaPostgresRepository) UpdateHarga(id int, tanggal string, harga float64) error {
	_, err := r.db.ExecContext(context.Background(),
		`UPDATE harga SET tanggal=$1, harga=$2 WHERE id=$3`, tanggal, harga, id)
	return err
}

func (r *HargaPostgresRepository) DeleteHarga(id int) error {
	_, err := r.db.ExecContext(context.Background(),
		`DELETE FROM harga WHERE id=$1`, id)
	return err
}

func (r *HargaPostgresRepository) CountHarga() (int, error) {
	var count int
	err := r.db.QueryRowContext(context.Background(),
		`SELECT COUNT(*) FROM harga`).Scan(&count)
	return count, err
}