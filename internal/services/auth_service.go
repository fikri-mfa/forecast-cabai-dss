package services

import (
	"errors"
	"forecast-cabai-dss/internal/auth"
	"forecast-cabai-dss/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserPostgresRepository
}

func NewAuthService(userRepo *repositories.UserPostgresRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(username, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return s.userRepo.CreateUser(username, string(hashed))
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("username atau password salah")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("username atau password salah")
	}
	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}