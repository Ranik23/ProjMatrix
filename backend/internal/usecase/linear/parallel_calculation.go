package linear

import (
	"ProjMatrix/pkg/wpool"
	"errors"
	"sync"
	"time"
)

func ParallelLinearFormCalculation(matrices [][][]float64, coefficients []float64, pool *wpool.WorkerPool) ([][]float64, float64, error) {
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

	var mu sync.Mutex

	for k, matrix := range matrices {
		coefficient := coefficients[k]

		// передаем задачу в пул воркеров
		pool.Submit(func() {
			partialResult := make([][]float64, rows)
			for i := range partialResult {
				partialResult[i] = make([]float64, cols)
			}

			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					partialResult[i][j] = coefficient * matrix[i][j]
				}
			}

			mu.Lock()
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					result[i][j] += partialResult[i][j]
				}
			}
			mu.Unlock()
		})
	}

	pool.Wait()

	elapsed := time.Since(start).Seconds()

	return result, elapsed, nil
}
