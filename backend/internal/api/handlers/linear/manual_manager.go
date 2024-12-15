package linear

import (
	"ProjMatrix/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleManualLinearForm(c *gin.Context, l *entity.LinearForm, workerClient entity.WorkersClient) error {
	log.Printf("Processing manual input of a linear form: %+v\n", *l)

	// это типо менджер такой именно для линейных форм, так как мы не сможем что-то что работало бы для всех типов

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
		err := manualNotParallel(c, l, worker.Client)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in processing calculations: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}
	case true:
		worker := workerClient.GetLeastLoadedWorker()
		worker.Valuation++
		err := manualParallel(c, l, worker.Client)
		worker.Valuation--
		if err != nil {
			log.Printf("Error in processing calculations: %w\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return err
		}

	}

	//matrices, err := mtrx.BuildMatrices(l.Matrices, l.MatrixSize.Rows, l.MatrixSize.Columns)
	//if err != nil {
	//	return fmt.Errorf("the matrix could not be formed: %w", err)
	//}
	//
	//data, err := json.Marshal(matrices)
	//if err != nil {
	//	return fmt.Errorf("the matrix could not be formed: %w", err)
	//}
	//
	//resp, err := workerClient.GetLinearFormCalculation(c, &proto.GetLinearFormCalculationRequest{
	//	Matrix:       string(data),
	//	Coefficients: l.Coefficients,
	//})
	//if err != nil {
	//	return fmt.Errorf("the matrix linear form could not be calculated: %w", err)
	//}

	//resultMatrix, timeCalc, err := linear.LinearFormCalculation(matrices, l.Coefficients)
	//if err != nil {
	//	return fmt.Errorf("the matrix linear form could not be calculated: %w", err)
	//}

	//pool := wpool.NewWorkerPool(runtime.NumCPU())
	//pool.Start()

	//_, parTimeCalc, err := linear.ParallelLinearFormCalculation(matrices, l.Coefficients, pool)
	//if err != nil {
	//	return fmt.Errorf("the matrix linear form could not be calculated in parallel: %w", err)
	//}
	//pool.Wait()
	//pool.Stop()

	//result := entity.CalculationResult{
	//	OperationType: l.OperationType,
	//	ResultMatrix:  converter.StringToMatrix(resp.ResultMatrix),
	//	TimeCalc:      resp.Time,
	//}
	//
	//// сохраняем результат в сессию
	//session.Set("calculationResult", result)
	//err = session.Save()
	//if err != nil {
	//	return fmt.Errorf("error saving the session: %w\n", err)
	//}
	//
	//c.Redirect(http.StatusFound, "/results")

	return nil
}
