package polynomial

import (
	"ProjMatrix/internal/entity"
	pol "ProjMatrix/internal/usecase/polynomial"
	mtrx "ProjMatrix/pkg/matrix"
	"ProjMatrix/pkg/wpool"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
)

func handleManualPolynomial(c *gin.Context, p *entity.Polynomial) error {
	log.Printf("Processing the manual input of a polynomial: %+v\n", *p)

	session := sessions.Default(c)

	matrix, err := mtrx.BuildMatrix(p.Matrix, p.MatrixSize.Rows, p.MatrixSize.Columns)
	if err != nil {
		return fmt.Errorf("the matrix could not be read: %w", err)
	}
	identityMatrix := mtrx.GenerateIdentityMatrix(p.MatrixSize.Rows)
	resultMatrix, timeCalc, err := pol.PolynomialCalculation(matrix, identityMatrix, p.Coefficients)
	if err != nil {
		return fmt.Errorf("the polynomial could not be calculated: %w", err)
	}

	pool := wpool.NewWorkerPool(runtime.NumCPU())
	pool.Start()

	_, par_timeCalc, err := pol.ParallelPolynomialCalculation(matrix, identityMatrix, p.Coefficients, pool)
	if err != nil {
		return fmt.Errorf("the polynomial could not be calculated in parallel: %w", err)
	}
	pool.Wait()
	pool.Stop()

	result := entity.CalculationResult{
		OperationType:    p.OperationType,
		ResultMatrix:     resultMatrix,
		TimeCalc:         timeCalc,
		TimeParallelCalc: par_timeCalc, // заглушка
	}

	session.Set("calculationResult", result)
	err = session.Save()
	if err != nil {
		return fmt.Errorf("error saving the session: %w\n", err)
	}

	c.Redirect(http.StatusFound, "/results")

	return nil
}
