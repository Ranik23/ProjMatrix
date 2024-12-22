package converter

import (
	"ProjMatrix/internal/entity"
	"encoding/json"
)

func ByteToMatrix(data []byte) [][]float64 {
	var matrices [][]float64
	_ = json.Unmarshal(data, &matrices)
	return matrices
}

func MatrixToByte(matrices [][]float64) []byte {
	data, _ := json.Marshal(matrices)
	return data

}

func CalculateResultFormToByte(data entity.CalculationResult) []byte {
	result, _ := json.Marshal(data)
	return result
}

func ByteToCalculationResult(data []byte) entity.CalculationResult {
	var result entity.CalculationResult
	_ = json.Unmarshal(data, &result)
	return result
}
