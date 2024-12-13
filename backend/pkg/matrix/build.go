package matrix

import "errors"

// BuildMatrix формирует матрицу, заданного размера, из массива
func BuildMatrix(array []float64, row, col int) ([][]float64, error) {
	if row <= 0 || col <= 0 {
		return nil, errors.New("the dimensions of the matrix must be greater than zero")
	}

	if len(array) != row*col {
		return nil, errors.New("the length of the array does not match the specified dimensions of the matrix")
	}

	matrix := make([][]float64, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]float64, col)
		for j := 0; j < col; j++ {
			matrix[i][j] = array[i*col+j] // Заполняем элементы матрицы
		}
	}

	return matrix, nil
}
