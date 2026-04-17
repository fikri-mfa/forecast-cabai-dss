package handlers

import (
	"encoding/json"
	"forecast-cabai-dss/internal/domain"
	"forecast-cabai-dss/internal/services"
	"net/http"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(as *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: as}
}
// Register godoc
// @Summary      Register user baru
// @Description  Mendaftarkan user baru ke sistem
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body domain.RegisterRequest true "Data registrasi"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request tidak valid", http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Password == "" {
		http.Error(w, "username dan password wajib diisi", http.StatusBadRequest)
		return
	}
	if err := h.authService.Register(req.Username, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "registrasi berhasil",
	})
}
// Login godoc
// @Summary      Login user
// @Description  Login dan mendapatkan JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body domain.LoginRequest true "Data login"
// @Success      200  {object}  domain.LoginResponse
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "request tidak valid", http.StatusBadRequest)
		return
	}
	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain.LoginResponse{Token: token})
}