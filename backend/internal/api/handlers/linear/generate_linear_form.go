package linear

import (
	"ProjMatrix/internal/entity"
	"ProjMatrix/internal/usecase/linear"
	mtrx "ProjMatrix/pkg/matrix"
	"ProjMatrix/pkg/wpool"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
)

func handleGeneratedLinearForm(c *gin.Context, l *entity.LinearForm) error {
	log.Printf("Обработка генерации линейной формы: %+v\n", *l)

	//parallel_result := ParallelPolynomialCalculation()
	matrices := make([][][]float64, l.MatrixCount)
	for i := 0; i < l.MatrixCount; i++ {
		matrix := mtrx.GenerateMatrix(l.MatrixSize.Rows, l.MatrixSize.Columns)
		matrices[i] = matrix
	}

	coefficients := mtrx.GenerateCoefficients(l.MatrixCount)

	_, timeCalc, err := linear.LinearFormCalculation(matrices, coefficients)
	if err != nil {
		return fmt.Errorf("не удалось вычислить матричную линейную форму: %w", err)
	}

	pool := wpool.NewWorkerPool(runtime.NumCPU())
	pool.Start()

	_, par_timeCalc, err := linear.ParallelLinearFormCalculation(matrices, coefficients, pool)
	if err != nil {
		return fmt.Errorf("не удалось вычислить матричную линейную форму: %w", err)
	}
	pool.Wait()
	pool.Stop()

	entity.ResultOfCalculations = entity.CalculationResult{
		OperationType:    l.OperationType,
		ResultMatrix:     nil,
		TimeCalc:         timeCalc,
		TimeParallelCalc: par_timeCalc,
	}

	c.Redirect(http.StatusFound, "/results")

	return nil
}
