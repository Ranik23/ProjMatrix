package polynomial

import (
	"ProjMatrix/internal/entity"
	pol "ProjMatrix/internal/usecase/polynomial"
	mtrx "ProjMatrix/pkg/matrix"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleManualPolynomial(c *gin.Context, p *entity.Polynomial) error {
	log.Printf("Обработка ручного ввода полинома: %+v\n", *p)

	//parallel_result := ParallelPolynomialCalculation()
	matrix, err := mtrx.BuildMatrix(p.Matrix, p.MatrixSize.Rows, p.MatrixSize.Columns)
	if err != nil {
		return fmt.Errorf("не удалось считать матрицу: %w", err)
	}
	identityMatrix := mtrx.GenerateIdentityMatrix(p.MatrixSize.Rows)
	resultMatrix, timeCalc, err := pol.PolynomialCalculation(matrix, identityMatrix, p.Coefficients)
	if err != nil {
		return fmt.Errorf("не удалось вычислить полином: %w", err)
	}

	entity.ResultOfCalculations = entity.CalculationResult{
		OperationType:    p.OperationType,
		ResultMatrix:     resultMatrix,
		TimeCalc:         timeCalc,
		TimeParallelCalc: 0.0, // заглушка
	}

	c.Redirect(http.StatusFound, "/results")

	return nil
}
