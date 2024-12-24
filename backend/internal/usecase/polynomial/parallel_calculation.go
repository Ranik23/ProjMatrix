package polynomial

import (
	"ProjMatrix/internal/converter"
	"ProjMatrix/internal/entity"
	"ProjMatrix/pkg/repository"
	"ProjMatrix/pkg/wpool"
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

func ParallelPolynomialCalculation(matrix, identityMatrix [][]float64, coefficients []float64, pool *wpool.WorkerPool, s repository.PgRepository) (string, float64, error) {
	start := time.Now()
	ctx := context.Background()
	if len(matrix) == 0 || len(identityMatrix) == 0 {
		return "", 0, errors.New("матрица или единичная матрица пустые")
	}
	if len(matrix) != len(matrix[0]) {
		return "", 0, errors.New("матрица должна быть квадратной")
	}

	n := len(matrix)

	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
	}

	// Добавляем первый коэффициент
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] += coefficients[0] * identityMatrix[i][j]
		}
	}

	// Временная матрица для хранения степеней матрицы
	currentPower := identityMatrix
	var err error

	for m := 1; m < len(coefficients); m++ {
		currentPower, err = ParallelMatrixMultiply(currentPower, matrix, pool)
		if err != nil {
			return "", 0, fmt.Errorf("не удалось вычислить Матричный Полином: %w", err)
		}

		pool.Submit(func() {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					result[i][j] += coefficients[m] * currentPower[i][j]
				}
			}
		})
	}

	pool.Wait()
	elapsed := time.Since(start).Seconds()

	session := entity.GenerateSessionID()
	err = s.Save(ctx, session, elapsed, converter.MatrixToByte(result))
	if err != nil {
		log.Println("Saving error")
	}
	return session, elapsed, nil
}

func ParallelMatrixMultiply(a, b [][]float64, pool *wpool.WorkerPool) ([][]float64, error) {
	if len(a) == 0 || len(b) == 0 || len(a[0]) != len(b) {
		return nil, errors.New("matrices are incompatible for multiplication")
	}

	rowsA, colsA := len(a), len(a[0])
	colsB := len(b[0])

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	var mu sync.Mutex
	for i := 0; i < rowsA; i++ {
		row := i
		pool.Submit(func() {
			for j := 0; j < colsB; j++ {
				sum := 0.0
				for k := 0; k < colsA; k++ {
					sum += a[row][k] * b[k][j]
				}
				mu.Lock()
				result[row][j] = sum
				mu.Unlock()
			}
		})
	}

	pool.Wait()
	return result, nil
}
