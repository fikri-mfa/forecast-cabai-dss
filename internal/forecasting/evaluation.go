package forecasting

import "math"

func MAPE(actual, forecast []float64) float64 {
	if len(actual) != len(forecast) {
		return 0
	}
	n := len(actual)
	sum := 0.0
	for i := 0; i < n; i++ {
		if actual[i] != 0 {
			sum += math.Abs((actual[i] - forecast[i]) / actual[i])
		}
	}
	return (sum / float64(n)) * 100
}

func RMSE(actual, forecast []float64) float64 {
	if len(actual) != len(forecast) {
		return 0
	}
	n := len(actual)
	sum := 0.0
	for i := 0; i < n; i++ {
		diff := actual[i] - forecast[i]
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(n))
}