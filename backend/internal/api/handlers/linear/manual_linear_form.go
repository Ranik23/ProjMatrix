package linear

import (
	"ProjMatrix/internal/entity"
	"ProjMatrix/internal/usecase/linear"
	mtrx "ProjMatrix/pkg/matrix"
	"ProjMatrix/pkg/wpool"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
)

func handleManualLinearForm(c *gin.Context, l *entity.LinearForm) error {
	log.Printf("Processing manual input of a linear form: %+v\n", *l)

	session := sessions.Default(c)

	matrices, err := mtrx.BuildMatrices(l.Matrices, l.MatrixSize.Rows, l.MatrixSize.Columns)
	if err != nil {
		return fmt.Errorf("the matrix could not be formed: %w", err)
	}
	resultMatrix, timeCalc, err := linear.LinearFormCalculation(matrices, l.Coefficients)
	if err != nil {
		return fmt.Errorf("the matrix linear form could not be calculated: %w", err)
	}

	pool := wpool.NewWorkerPool(runtime.NumCPU())
	pool.Start()

	_, parTimeCalc, err := linear.ParallelLinearFormCalculation(matrices, l.Coefficients, pool)
	if err != nil {
		return fmt.Errorf("the matrix linear form could not be calculated in parallel: %w", err)
	}
	pool.Wait()
	pool.Stop()

	result := entity.CalculationResult{
		OperationType:    l.OperationType,
		ResultMatrix:     resultMatrix,
		TimeCalc:         timeCalc,
		TimeParallelCalc: parTimeCalc,
	}

	// сохраняем результат в сессию
	session.Set("calculationResult", result)
	err = session.Save()
	if err != nil {
		return fmt.Errorf("error saving the session: %w\n", err)
	}

	c.Redirect(http.StatusFound, "/results")

	return nil
}
