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

func handleManualLinearForm(c *gin.Context, l *entity.LinearForm) error {
	log.Printf("Обработка ручного ввода линейной формы: %+v\n", *l)

	//parallel_result := ParallelPolynomialCalculation()
	matrices, err := mtrx.BuildMatrices(l.Matrices, l.MatrixSize.Rows, l.MatrixSize.Columns)
	if err != nil {
		return fmt.Errorf("не сформировать матрицу: %w", err)
	}
	resultMatrix, timeCalc, err := linear.LinearFormCalculation(matrices, l.Coefficients)
	if err != nil {
		return fmt.Errorf("не удалось вычислить матричную линейную форму: %w", err)
	}

	pool := wpool.NewWorkerPool(runtime.NumCPU())
	pool.Start()

	_, par_timeCalc, err := linear.ParallelLinearFormCalculation(matrices, l.Coefficients, pool)
	if err != nil {
		return fmt.Errorf("не удалось вычислить матричную линейную форму: %w", err)
	}
	pool.Wait()
	pool.Stop()

	entity.ResultOfCalculations = entity.CalculationResult{
		OperationType:    l.OperationType,
		ResultMatrix:     resultMatrix,
		TimeCalc:         timeCalc,
		TimeParallelCalc: par_timeCalc,
	}

	c.Redirect(http.StatusFound, "/results")

	return nil
}
