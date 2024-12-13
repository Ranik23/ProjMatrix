package matrix

import "errors"

func MatrixMultiply(a, b [][]float64) ([][]float64, error) {
	if len(a) == 0 || len(b) == 0 || len(a[0]) != len(b) {
		return nil, errors.New("matrices are incompatible for multiplication")
	}

	rowsA, colsA := len(a), len(a[0])
	colsB := len(b[0])

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return result, nil
}
