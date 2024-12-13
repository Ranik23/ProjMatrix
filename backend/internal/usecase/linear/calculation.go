package linear

import (
	"errors"
	"time"
)

func LinearFormCalculation(matrices [][][]float64, coefficients []float64) ([][]float64, float64, error) {
	start := time.Now()

	if len(matrices) == 0 || len(coefficients) == 0 {
		return nil, 0, errors.New("the array of matrices or the array of coefficients is empty")
	}

	if len(matrices) != len(coefficients) {
		return nil, 0, errors.New("the number of matrices and coefficients must match")
	}

	rows, cols := len(matrices[0]), len(matrices[0][0])

	result := make([][]float64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
	}

	for k, matrix := range matrices {
		coefficient := coefficients[k]
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				result[i][j] += coefficient * matrix[i][j]
			}
		}
	}

	elapsed := time.Since(start).Seconds()

	return result, elapsed, nil
}
