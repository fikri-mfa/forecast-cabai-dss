package forecasting

import "math"

type OptimalParams struct {
	Alpha  float64
	Beta   float64
	Gamma  float64
	MAPE   float64
}

func FindOptimalParams(data []float64, seasonLength int) OptimalParams {
	best := OptimalParams{
		Alpha: 0.1,
		Beta:  0.1,
		Gamma: 0.1,
		MAPE:  math.MaxFloat64,
	}

	steps := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9}

	for _, a := range steps {
		for _, b := range steps {
			for _, g := range steps {
				result := TripleExponentialSmoothing(data, a, b, g, seasonLength, 1)
				mape := MAPE(data, result.Fitted)
				if mape < best.MAPE {
					best = OptimalParams{Alpha: a, Beta: b, Gamma: g, MAPE: mape}
				}
			}
		}
	}

	return best
}