package matrix

import (
	"math/rand"
	"time"
)

func GenerateCoefficients(length int) []float64 {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	coefficients := make([]float64, length)
	for i := 0; i < length; i++ {
		coefficients[i] = rng.Float64()*2000 - 1000
	}

	return coefficients
}
