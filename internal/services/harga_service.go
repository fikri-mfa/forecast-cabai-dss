package services

import (
	"forecast-cabai-dss/internal/domain"
	"forecast-cabai-dss/internal/repositories"
)

type HargaService struct {
	hargaRepo *repositories.HargaPostgresRepository
}

func NewHargaService(repo *repositories.HargaPostgresRepository) *HargaService {
	return &HargaService{hargaRepo: repo}
}

func (s *HargaService) GetAllHarga() ([]domain.HargaRow, error) {
	return s.hargaRepo.GetAllHargaWithDate()
}

func (s *HargaService) CreateHarga(tanggal string, harga float64) error {
	return s.hargaRepo.CreateHarga(tanggal, harga)
}

func (s *HargaService) UpdateHarga(id int, tanggal string, harga float64) error {
	return s.hargaRepo.UpdateHarga(id, tanggal, harga)
}

func (s *HargaService) DeleteHarga(id int) error {
	return s.hargaRepo.DeleteHarga(id)
}