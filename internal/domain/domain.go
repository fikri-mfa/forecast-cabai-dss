package domain

type ParameterTES struct {
	Alpha        float64
	Beta         float64
	Gamma        float64
	SeasonLength int
}

type User struct {
	ID       int
	Username string
	Password string
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}