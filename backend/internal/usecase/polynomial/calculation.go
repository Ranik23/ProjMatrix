package polynomial

import (
	"ProjMatrix/internal/converter"
	"ProjMatrix/internal/entity"
	mtrx "ProjMatrix/pkg/matrix"
	"ProjMatrix/pkg/repository"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func PolynomialCalculation(matrix, identityMatrix [][]float64, coefficients []float64, s repository.PgRepository) (string, float64, error) {
	start := time.Now()
	ctx := context.Background()
	if len(matrix) == 0 || len(identityMatrix) == 0 {
		return "", 0, errors.New("the matrix or the unit matrix is empty")
	}
	if len(matrix) != len(matrix[0]) {
		return "", 0, errors.New("the matrix must be square")
	}

	n := len(matrix)

	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] += coefficients[0] * identityMatrix[i][j]
		}
	}

	// Временная матрица для хранения степеней матрицы A
	currentPower := identityMatrix
	var err error

	for m := 1; m < len(coefficients); m++ {

		currentPower, err = mtrx.MatrixMultiply(currentPower, matrix)
		if err != nil {
			return "", 0, fmt.Errorf("the Matrix Polynomial could not be calculated: %w", err)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				result[i][j] += coefficients[m] * currentPower[i][j]
			}
		}
	}

	elapsed := time.Since(start).Seconds()

	session := entity.GenerateSessionID()
	err = s.Save(ctx, session, elapsed, converter.MatrixToByte(result))
	if err != nil {
		log.Println("Saving error")
	}

	return session, elapsed, nil
}
