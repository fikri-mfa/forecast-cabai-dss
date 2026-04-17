package repositories

type HargaRepository interface {
	GetAllHarga() ([]float64, error)
}
