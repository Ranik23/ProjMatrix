package linear

import (
	"ProjMatrix/internal/entity"
	"ProjMatrix/pkg/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func manualParallel(c *gin.Context, l *entity.LinearForm, workerClient proto.WorkerServiceClient) error {

	//session := sessions.Default(c)
	//
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
	//resp, err := workerClient.GetParallelLinearFormCalculation(c, &proto.GetParallelLinearFormCalculationRequest{
	//	Matrix:       data,
	//	Coefficients: l.Coefficients,
	//})
	//if err != nil {
	//	return fmt.Errorf("the matrix linear form could not be calculated: %w", err)
	//}
	//
	//result := entity.CalculationResult{
	//	OperationType: l.OperationType,
	//	ResultMatrix:  converter.ByteToMatrix(resp.ResultMatrix),
	//	TimeCalc:      resp.Time,
	//}
	//
	//// сохраняем результат в сессию
	//session.Set("calculationResult", result)
	//err = session.Save()
	//if err != nil {
	//	return fmt.Errorf("error saving the session: %w\n", err)
	//}

	c.Redirect(http.StatusFound, "/results")

	return nil
}
