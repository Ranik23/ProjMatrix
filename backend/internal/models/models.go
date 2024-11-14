package models

type Matrix struct {
	Rows int         `json:"rows"`
	Cols int         `json:"cols"`
	Data [][]float64 `json:"data"`
}

type Polynomial struct {
	Coefficients []float64 `json:"coefficients"`
	Degree       int       `json:"degree"`
}

type LinearForm struct {
	Coefficients []float64 `json:"coefficients"`
	Matrices     []Matrix  `json:"matrices"`
	Count        int       `json:"count"`
}
