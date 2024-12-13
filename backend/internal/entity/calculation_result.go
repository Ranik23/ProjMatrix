package entity

type CalculationResult struct {
	OperationType    string      `json:"operationType"`    // Тип операции
	ResultMatrix     [][]float64 `json:"resultMatrix"`     // Результирующая матрица
	TimeCalc         float64     `json:"timeCalc"`         // Время последовательных вычислений (в секундах)
	TimeParallelCalc float64     `json:"timeParallelCalc"` // Время параллельных вычислений (в секундах)
}
