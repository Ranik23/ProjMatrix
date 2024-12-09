package matrix

import "errors"

func BuildMatrix(array []float64, row, col int) ([][]float64, error) {
	if row <= 0 || col <= 0 {
		return nil, errors.New("размеры матрицы должны быть больше нуля")
	}

	if len(array) != row*col {
		return nil, errors.New("длина массива не соответствует заданным размерам матрицы")
	}

	// Создаем матрицу
	matrix := make([][]float64, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]float64, col)
		for j := 0; j < col; j++ {
			matrix[i][j] = array[i*col+j] // Заполняем элементы матрицы
		}
	}

	return matrix, nil
}
