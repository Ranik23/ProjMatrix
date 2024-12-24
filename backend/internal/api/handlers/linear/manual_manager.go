package linear

import (
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleManualLinearForm(c *gin.Context, l *entity.LinearForm, workerClient *entity.WorkersClient) error {
	log.Printf("Processing manual input of a linear form: %+v\n", *l)

	taskSize := l.MatrixSize.Rows * l.MatrixSize.Columns
	totalOps := taskSize * l.MatrixCount
	needParallel := false

	switch {
	case (taskSize > 2500 && taskSize <= 10000) && l.MatrixCount > 10:
		needParallel = true
	case taskSize > 10000 || l.MatrixCount > 50:
		needParallel = true
	}

	if totalOps > 100000 {
		needParallel = true
	}

	switch needParallel {
	case false:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := manualNotParallel(c, l, worker.Client, workerClient)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in processing calculations: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}
	case true:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := manualParallel(c, l, worker.Client, workerClient)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in processing calculations: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}

	}
	return nil
}
