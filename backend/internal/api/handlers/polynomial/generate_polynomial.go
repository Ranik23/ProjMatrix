package polynomial

import (
	"ProjMatrix/internal/entity"
	pol "ProjMatrix/internal/usecase/polynomial"
	mtrx "ProjMatrix/pkg/matrix"
	"ProjMatrix/pkg/wpool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
)

func handleGeneratedPolynomial(c *gin.Context, p *entity.Polynomial) error {
	log.Printf("Обработка генерации полинома: %+v\n", *p)

	matrix := mtrx.GenerateMatrix(p.MatrixSize.Rows, p.MatrixSize.Columns)
	coefficients := mtrx.GenerateCoefficients(p.Degree)

	identityMatrix := mtrx.GenerateIdentityMatrix(p.MatrixSize.Rows)
	_, timeCalc, err := pol.PolynomialCalculation(matrix, identityMatrix, coefficients)
	if err != nil {
		return fmt.Errorf("не удалось вычислить полином: %w", err)
	}

	pool := wpool.NewWorkerPool(runtime.NumCPU())
	pool.Start()
	_, par_timeCalc, err := pol.ParallelPolynomialCalculation(matrix, identityMatrix, coefficients, pool)
	if err != nil {
		return fmt.Errorf("не удалось вычислить полином: %w", err)
	}
	pool.Wait()
	pool.Stop()

	entity.ResultOfCalculations = entity.CalculationResult{
		OperationType:    p.OperationType,
		ResultMatrix:     nil,
		TimeCalc:         timeCalc,
		TimeParallelCalc: par_timeCalc, // заглушка
	}

	c.Redirect(http.StatusFound, "/results")

	return nil
}