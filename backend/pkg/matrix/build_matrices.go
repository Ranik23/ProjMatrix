package matrix

import "fmt"

// BuildMatrices формирует массив матриц из массива массивов
func BuildMatrices(arrays [][]float64, rows, cols int) ([][][]float64, error) {
	var matrices [][][]float64

	for _, array := range arrays {
		matrix, err := BuildMatrix(array, rows, cols)
		if err != nil {
			return nil, fmt.Errorf("ошибка при создании матрицы: %w", err)
		}
		matrices = append(matrices, matrix)
	}

	return matrices, nil
}
