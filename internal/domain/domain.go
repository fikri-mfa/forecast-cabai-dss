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
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"password123"`
}

type LoginRequest struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"password123"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type Forecast struct {
	ID           int       `json:"id" example:"1"`
	UserID       int       `json:"user_id" example:"1"`
	Alpha        float64   `json:"alpha" example:"0.2"`
	Beta         float64   `json:"beta" example:"0.1"`
	Gamma        float64   `json:"gamma" example:"0.1"`
	SeasonLength int       `json:"season_length" example:"12"`
	Periods      int       `json:"periods" example:"3"`
	Result       []float64 `json:"result"`
	CreatedAt    string    `json:"created_at" example:"2024-01-15T10:30:00Z"`
}

type ForecastRequest struct {
	Alpha        float64 `json:"alpha" example:"0.2"`
	Beta         float64 `json:"beta" example:"0.1"`
	Gamma        float64 `json:"gamma" example:"0.1"`
	SeasonLength int     `json:"season_length" example:"12"`
	Periods      int     `json:"periods" example:"3"`
}

type EvaluationResult struct {
	MAPE float64 `json:"mape" example:"5.23"`
	RMSE float64 `json:"rmse" example:"1250.75"`
}

type ForecastResponse struct {
	Forecast   []float64        `json:"forecast"`
	Evaluation EvaluationResult `json:"evaluation"`
}