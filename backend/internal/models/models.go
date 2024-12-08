package models

type MatrixSize struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
}

type Polynomial struct {
	OperationType string     `json:"operationType"`
	MatrixSize    MatrixSize `json:"matrixSize"`
	Matrix        []float64  `json:"matrix,omitempty"`
	Degree        int        `json:"degree"`
	Coefficients  []float64  `json:"coefficients,omitempty"`
}

type LinearForm struct {
	OperationType string      `json:"operationType"`
	MatrixCount   int         `json:"matrixCount"`
	MatrixSize    MatrixSize  `json:"matrixSize"`
	Matrices      [][]float64 `json:"matrices,omitempty"`
	Coefficients  []float64   `json:"coefficients,omitempty"`
}
