package polynomial

import (
	"ProjMatrix/internal/entity"
	pol "ProjMatrix/internal/usecase/polynomial"
	mtrx "ProjMatrix/pkg/matrix"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func handleGeneratedPolynomial(c *gin.Context, p *entity.Polynomial) error {
	log.Printf("Обработка генерации полинома: %+v\n", *p)

	//parallel_result := ParallelPolynomialCalculation()
	matrix := mtrx.GenerateMatrix(p.MatrixSize.Rows, p.MatrixSize.Columns)
	coefficients := GenerateCoefficients(p.Degree)

	identityMatrix := mtrx.GenerateIdentityMatrix(p.MatrixSize.Rows)
	_, timeCalc, err := pol.PolynomialCalculation(matrix, identityMatrix, coefficients)
	if err != nil {
		return fmt.Errorf("не удалось вычислить полином: %w", err)
	}
	
	entity.ResultOfCalculations = entity.CalculationResult{
		OperationType:    p.OperationType,
		ResultMatrix:     nil,
		TimeCalc:         timeCalc,
		TimeParallelCalc: 0.0, // заглушка
	}

	c.Redirect(http.StatusFound, "/results")

	return nil
}

func GenerateCoefficients(length int) []float64 {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	coefficients := make([]float64, length)
	for i := 0; i < length; i++ {
		coefficients[i] = rng.Float64()*2000 - 1000
	}

	return coefficients
}
