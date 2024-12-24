package polynomial

import (
	"ProjMatrix/internal/entity"
	mtrx "ProjMatrix/pkg/matrix"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleGeneratedPolynomial(c *gin.Context, p *entity.Polynomial, workerClient *entity.WorkersClient) error {
	log.Printf("Processing polynomial generation: %+v\n", *p)

	matrix := mtrx.GenerateMatrix(p.MatrixSize.Rows, p.MatrixSize.Columns)
	coefficients := mtrx.GenerateCoefficients(p.Degree)
	identityMatrix := mtrx.GenerateIdentityMatrix(p.MatrixSize.Rows)

	taskSize := p.MatrixSize.Rows * p.MatrixSize.Columns
	totalOps := taskSize * len(coefficients)
	needParallel := false

	switch {
	case (taskSize > 2500 && taskSize <= 10000) && len(coefficients) > 10:
		needParallel = true
	case taskSize > 10000 || len(coefficients) > 50:
		needParallel = true
	}

	if totalOps > 100000 {
		needParallel = true
	}

	switch needParallel {
	case false:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := generateNotParallel(c, matrix, coefficients, identityMatrix, worker.Client, workerClient)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in sequential polynomial generation: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}
	case true:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := generateParallel(c, matrix, coefficients, identityMatrix, worker.Client, workerClient)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in parallel polynomial generation: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}
	}

	return nil
}
