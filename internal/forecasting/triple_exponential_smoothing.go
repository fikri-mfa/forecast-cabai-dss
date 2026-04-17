package forecasting

type TESResult struct {
	Forecast []float64
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
	seasonal := make([]float64, n)
	forecast := make([]float64, forecastPeriods)

	// 1. Inisialisasi awal
	level[0] = data[0]
	trend[0] = data[1] - data[0]

	for i := 0; i < seasonLength; i++ {
		seasonal[i] = data[i] - level[0]
	}

	// 2. Iterasi data historis
	for i := 1; i < n; i++ {
		seasonIndex := i % seasonLength

		level[i] = alpha*(data[i]-seasonal[seasonIndex]) +
			(1-alpha)*(level[i-1]+trend[i-1])

		trend[i] = beta*(level[i]-level[i-1]) +
			(1-beta)*trend[i-1]

		seasonal[seasonIndex] = gamma*(data[i]-level[i]) +
			(1-gamma)*seasonal[seasonIndex]
	}

	// 3. Forecast ke depan
	lastLevel := level[n-1]
	lastTrend := trend[n-1]

	for i := 1; i <= forecastPeriods; i++ {
		seasonIndex := (n + i - 1) % seasonLength
		forecast[i-1] = lastLevel +
			float64(i)*lastTrend +
			seasonal[seasonIndex]
	}

	return TESResult{
		Forecast: forecast,
		Level:    level,
		Trend:    trend,
		Seasonal: seasonal,
	}
}
