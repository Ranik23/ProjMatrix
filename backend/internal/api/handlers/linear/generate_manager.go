package linear

import (
	"ProjMatrix/internal/entity"
	mtrx "ProjMatrix/pkg/matrix"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleGeneratedLinearForm(c *gin.Context, l *entity.LinearForm, workerClient entity.WorkersClient) error {
	log.Printf("Processing of linear shape generation: %+v\n", *l)

	matrices := make([][][]float64, l.MatrixCount)
	for i := 0; i < l.MatrixCount; i++ {
		matrix := mtrx.GenerateMatrix(l.MatrixSize.Rows, l.MatrixSize.Columns)
		matrices[i] = matrix
	}

	coefficients := mtrx.GenerateCoefficients(l.MatrixCount)

	taskSize := l.MatrixSize.Rows * l.MatrixSize.Columns
	totalOps := taskSize * l.MatrixCount
	needParallel := false

	switch {
	case (taskSize > 2500 && taskSize <= 10000) && l.MatrixCount > 10:
		// Промежуточный случай: средний размер матриц и их количество большое
		needParallel = true
	case taskSize > 10000 || l.MatrixCount > 50:
		// Крупные матрицы или очень большое количество — параллелим
		needParallel = true
	}

	// Дополнительная проверка на общее количество операций
	if totalOps > 100000 {
		needParallel = true
	}

	switch needParallel {
	case false:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := generateNotParallel(c, matrices, coefficients, worker.Client)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in processing calculations: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}
	case true:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := generateParallel(c, matrices, coefficients, worker.Client)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in processing calculations: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}

	}

	return nil
}
