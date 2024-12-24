package linear

import (
	"ProjMatrix/internal/converter"
	"ProjMatrix/internal/entity"
	"ProjMatrix/pkg/repository"
	"context"
	"errors"
	"log"
	"time"
)

func LinearFormCalculation(matrices [][][]float64, coefficients []float64, s repository.PgRepository) (string, float64, error) {
	start := time.Now()
	ctx := context.Background()
	if len(matrices) == 0 || len(coefficients) == 0 {
		return "", 0, errors.New("the array of matrices or the array of coefficients is empty")
	}

	if len(matrices) != len(coefficients) {
		return "", 0, errors.New("the number of matrices and coefficients must match")
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

	session := entity.GenerateSessionID()
	err := s.Save(ctx, session, elapsed, converter.MatrixToByte(result))
	if err != nil {
		log.Println("Saving error")
	}

	return session, elapsed, nil
}
