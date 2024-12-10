package matrix

import (
	"math/rand"
	"time"
)

func GenerateMatrix(rows, cols int) [][]float64 {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = rng.Float64()*2000 - 1000 // генерация от -1000 до 1000
		}
	}

	return matrix
}
