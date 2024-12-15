package linear

import (
	"ProjMatrix/internal/converter"
	"ProjMatrix/internal/entity"
	"ProjMatrix/pkg/proto"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func generateNotParallel(c *gin.Context, matrices [][][]float64, coefficients []float64, workerClient proto.WorkerServiceClient) error {
	session := sessions.Default(c)

	data, err := json.Marshal(matrices)
	if err != nil {
		return fmt.Errorf("the matrix could not be formed: %w", err)
	}

	resp, err := workerClient.GetLinearFormCalculation(c, &proto.GetLinearFormCalculationRequest{
		Matrix:       data,
		Coefficients: coefficients,
	})
	if err != nil {
		return fmt.Errorf("the matrix linear form could not be calculated: %w", err)
	}

	result := entity.CalculationResult{
		OperationType: resp.OperationType,
		ResultMatrix:  converter.ByteToMatrix(resp.ResultMatrix),
		TimeCalc:      resp.Time,
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
