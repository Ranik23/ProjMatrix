package entity

type LinearForm struct {
	OperationType string      `json:"operationType"`
	MatrixCount   int         `json:"matrixCount"`
	MatrixSize    MatrixSize  `json:"matrixSize"`
	Matrices      [][]float64 `json:"matrices,omitempty"`
	Coefficients  []float64   `json:"coefficients,omitempty"`
}
