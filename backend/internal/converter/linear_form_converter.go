package converter

import "encoding/json"

func ByteToMatrix(data []byte) [][]float64 {
	var matrices [][]float64
	_ = json.Unmarshal(data, &matrices)
	return matrices
}
