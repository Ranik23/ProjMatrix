package entity

type Polynomial struct {
	OperationType string     `json:"operationType"`
	MatrixSize    MatrixSize `json:"matrixSize"`
	Matrix        []float64  `json:"matrix,omitempty"`
	Degree        int        `json:"degree"`
	Coefficients  []float64  `json:"coefficients,omitempty"`
}
