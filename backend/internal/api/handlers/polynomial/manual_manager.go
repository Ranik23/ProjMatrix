package polynomial

import (
	"ProjMatrix/internal/entity"
	mtrx "ProjMatrix/pkg/matrix"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleManualPolynomial(c *gin.Context, p *entity.Polynomial, workerClient *entity.WorkersClient) error {
	log.Printf("Processing the manual input of a polynomial: %+v\n", *p)

	matrix, err := mtrx.BuildMatrix(p.Matrix, p.MatrixSize.Rows, p.MatrixSize.Columns)
	if err != nil {
		return fmt.Errorf("the matrix could not be read: %w", err)
	}
	identityMatrix := mtrx.GenerateIdentityMatrix(p.MatrixSize.Rows)
	taskSize := p.MatrixSize.Rows * p.MatrixSize.Columns
	totalOps := taskSize * len(p.Coefficients)
	needParallel := false

	switch {
	case (taskSize > 2500 && taskSize <= 10000) && len(p.Coefficients) > 10:
		needParallel = true
	case taskSize > 10000 || len(p.Coefficients) > 50:
		needParallel = true
	}

	if totalOps > 100000 {
		needParallel = true
	}

	switch needParallel {
	case false:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := manualNotParallel(c, matrix, identityMatrix, p.Coefficients, worker.Client, workerClient)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in sequential polynomial calculation: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}
	case true:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := manualParallel(c, matrix, identityMatrix, p.Coefficients, worker.Client, workerClient)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in parallel polynomial calculation: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}
	}

	return nil
}
