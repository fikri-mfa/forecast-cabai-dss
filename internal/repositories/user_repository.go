package repositories

import (
	"database/sql"
	"errors"
	"forecast-cabai-dss/internal/domain"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) CreateUser(username, hashedPassword string) error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2)`
	_, err := r.db.Exec(query, username, hashedPassword)
	return err
}

func (r *UserPostgresRepository) GetUserByUsername(username string) (*domain.User, error) {
	query := `SELECT id, username, password FROM users WHERE username = $1`
	row := r.db.QueryRow(query, username)

	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}