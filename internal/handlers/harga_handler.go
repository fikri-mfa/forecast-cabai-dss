package handlers

import (
	"encoding/json"
	"forecast-cabai-dss/internal/domain"
	"forecast-cabai-dss/internal/services"
	"net/http"
	"strconv"
	"strings"
)

type HargaHandler struct {
	hargaService *services.HargaService
}

func NewHargaHandler(hs *services.HargaService) *HargaHandler {
	return &HargaHandler{hargaService: hs}
}

// GET /harga
func (h *HargaHandler) GetAllHarga(w http.ResponseWriter, r *http.Request) {
	data, err := h.hargaService.GetAllHarga()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}

// POST /harga
func (h *HargaHandler) CreateHarga(w http.ResponseWriter, r *http.Request) {
	var req domain.HargaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request tidak valid", http.StatusBadRequest)
		return
	}
	if err := h.hargaService.CreateHarga(req.Tanggal, req.Harga); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// PUT /harga/{id}
func (h *HargaHandler) UpdateHarga(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/harga/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id tidak valid", http.StatusBadRequest)
		return
	}
	var req domain.HargaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request tidak valid", http.StatusBadRequest)
		return
	}
	if err := h.hargaService.UpdateHarga(id, req.Tanggal, req.Harga); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// DELETE /harga/{id}
func (h *HargaHandler) DeleteHarga(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/harga/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id tidak valid", http.StatusBadRequest)
		return
	}
	if err := h.hargaService.DeleteHarga(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}