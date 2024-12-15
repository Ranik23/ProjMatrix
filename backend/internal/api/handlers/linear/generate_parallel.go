package linear

import (
	"ProjMatrix/internal/converter"
	"ProjMatrix/internal/entity"
	"ProjMatrix/pkg/proto"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func generateParallel(c *gin.Context, matrices [][][]float64, coefficients []float64, workerClient proto.WorkerServiceClient) error {
	session := sessions.Default(c)

	stream, err := workerClient.GetParallelLinearFormCalculation(c)
	if err != nil {
		return fmt.Errorf("error opening stream: %w", err)
	}

	// Отправка данных через стрим
	data, err := json.Marshal(matrices)
	if err != nil {
		return fmt.Errorf("the matrix could not be formed: %w", err)
	}

	err = stream.Send(&proto.GetParallelLinearFormCalculationRequest{
		Matrix:       data,
		Coefficients: coefficients,
	})
	if err != nil {
		return fmt.Errorf("error sending data to stream: %w", err)
	}

	stream.CloseSend()

	// Получение данных из стрима
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error receiving stream data: %w", err)
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
	}

	c.Redirect(http.StatusFound, "/results")
	return nil
}
