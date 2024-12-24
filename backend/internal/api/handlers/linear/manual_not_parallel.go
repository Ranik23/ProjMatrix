package linear

import (
	"ProjMatrix/internal/entity"
	"ProjMatrix/pkg/proto"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func manualNotParallel(c *gin.Context, l *entity.LinearForm, workerClient proto.WorkerServiceClient, client *entity.WorkersClient) error {
	log.Printf("Starting parallel calculation with matrix size: %d", len(l.Matrices))

	stream, err := workerClient.GetLinearFormCalculation(c)
	if err != nil {
		return fmt.Errorf("error opening stream: %w", err)
	}

	data := struct {
		Matrices     [][][]float64 `json:"matrices"`
		Coefficients []float64     `json:"coefficients"`
	}{
		Matrices:     [][][]float64{l.Matrices},
		Coefficients: l.Coefficients,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	const chunkSize = 1 * 1024 * 1024
	totalChunks := (len(jsonData) + chunkSize - 1) / chunkSize
	log.Printf("Total data size: %d bytes, will be sent in %d chunks", len(jsonData), totalChunks)

	for i := 0; i < len(jsonData); i += chunkSize {
		end := i + chunkSize
		if end > len(jsonData) {
			end = len(jsonData)
		}

		chunk := &proto.Chunk{
			Content: jsonData[i:end],
		}

		log.Printf("Sending chunk %d/%d, size: %d bytes", (i/chunkSize)+1, totalChunks, len(chunk.Content))

		if err := stream.Send(chunk); err != nil {
			return fmt.Errorf("error sending chunk %d/%d: %w", (i/chunkSize)+1, totalChunks, err)
		}
	}

	log.Println("All chunks sent, closing send stream")
	if err := stream.CloseSend(); err != nil {
		return fmt.Errorf("error closing send stream: %w", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		if err == io.EOF {
			return fmt.Errorf("stream closed without receiving response")
		}
		return fmt.Errorf("error receiving response: %w", err)
	}

	result := entity.CalculationResult{
		OperationType: "Parallel Calc Lin Form",
		TimeCalc:      resp.Time,
	}

	client.Session = resp.Key

	log.Printf("Received calculation result, operation: %s, time: %f",
		result.OperationType, result.TimeCalc)

	c.Redirect(http.StatusFound, "/results")
	return nil
}
