package forecasting

type TESResult struct {
	Forecast []float64
	Fitted   []float64
	Level    []float64
	Trend    []float64
	Seasonal []float64
}

func TripleExponentialSmoothing(
	data []float64,
	alpha, beta, gamma float64,
	seasonLength int,
	forecastPeriods int,
) TESResult {
	n := len(data)
	level := make([]float64, n)
	trend := make([]float64, n)
	seasonal := make([]float64, n+forecastPeriods)
	fitted := make([]float64, n)
	forecast := make([]float64, forecastPeriods)

	// 1. Inisialisasi Level — rata-rata periode pertama
	sum1 := 0.0
	for i := 0; i < seasonLength; i++ {
		sum1 += data[i]
	}
	L := sum1 / float64(seasonLength)

	// Inisialisasi Level periode kedua — untuk hitung Trend
	sum2 := 0.0
	for i := seasonLength; i < 2*seasonLength; i++ {
		sum2 += data[i]
	}
	L2 := sum2 / float64(seasonLength)

	// Inisialisasi Trend
	T := (L2 - L) / float64(seasonLength)

	// 2. Inisialisasi Seasonal — multiplicative: S[i] = harga[i] / L
	S := make([]float64, seasonLength)
	for i := 0; i < seasonLength; i++ {
		S[i] = data[i] / L
	}

	// 3. Iterasi semua data
	for t := 0; t < n; t++ {
		i := t % seasonLength
		if t < seasonLength {
			// Periode awal — belum ada forecast
			fitted[t] = 0
			level[t] = L
			trend[t] = T
			seasonal[t] = S[i]
		} else {
			// Forecast = (L + T) * S
			fitted[t] = (L + T) * S[i]

			// Update Level, Trend, Seasonal
			L_prev := L
			L = alpha*(data[t]/S[i]) + (1-alpha)*(L+T)
			T = beta*(L-L_prev) + (1-beta)*T
			S[i] = gamma*(data[t]/L) + (1-gamma)*S[i]

			level[t] = L
			trend[t] = T
			seasonal[t] = S[i]
		}
	}

	// 4. Forecast ke depan
	for i := 1; i <= forecastPeriods; i++ {
		seasonIndex := (n + i - 1) % seasonLength
		forecast[i-1] = (L + float64(i)*T) * S[seasonIndex]
	}

	return TESResult{
		Forecast: forecast,
		Fitted:   fitted,
		Level:    level,
		Trend:    trend,
		Seasonal: seasonal[:n],
	}
}