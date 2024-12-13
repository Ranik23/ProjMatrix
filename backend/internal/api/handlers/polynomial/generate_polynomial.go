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

func handleGeneratedPolynomial(c *gin.Context, p *entity.Polynomial) error {
	log.Printf("Processing of polynomial generation: %+v\n", *p)

	session := sessions.Default(c)

	matrix := mtrx.GenerateMatrix(p.MatrixSize.Rows, p.MatrixSize.Columns)
	coefficients := mtrx.GenerateCoefficients(p.Degree)

	identityMatrix := mtrx.GenerateIdentityMatrix(p.MatrixSize.Rows)
	_, timeCalc, err := pol.PolynomialCalculation(matrix, identityMatrix, coefficients)
	if err != nil {
		return fmt.Errorf("the polynomial could not be calculated: %w", err)
	}

	pool := wpool.NewWorkerPool(runtime.NumCPU())
	pool.Start()
	_, parTimeCalc, err := pol.ParallelPolynomialCalculation(matrix, identityMatrix, coefficients, pool)
	if err != nil {
		return fmt.Errorf("the polynomial could not be calculated in parallel: %w", err)
	}
	pool.Wait()
	pool.Stop()

	result := entity.CalculationResult{
		OperationType:    p.OperationType,
		ResultMatrix:     nil,
		TimeCalc:         timeCalc,
		TimeParallelCalc: parTimeCalc, // заглушка
	}

	session.Set("calculationResult", result)
	err = session.Save()
	if err != nil {
		return fmt.Errorf("error saving the session: %w\n", err)
	}

	c.Redirect(http.StatusFound, "/results")

	return nil
}
