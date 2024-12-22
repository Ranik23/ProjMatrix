package worker

import (
	"ProjMatrix/internal/usecase/linear"
	desc "ProjMatrix/pkg/proto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func (s *Service) GetParallelLinearFormCalculation(stream desc.WorkerService_GetParallelLinearFormCalculationServer) error {
	log.Printf("GetParallelLinearFormCalculation method was invoked")

	var buffer bytes.Buffer

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Finished receiving chunks, total size: %d bytes", buffer.Len())
			break
		}
		if err != nil {
			log.Printf("Error receiving stream: %v", err)
			return err
		}

		if _, err := buffer.Write(chunk.Content); err != nil {
			log.Printf("Error writing to buffer: %v", err)
			return err
		}

		log.Printf("Received chunk of size: %d bytes", len(chunk.Content))
	}

	if buffer.Len() == 0 {
		return fmt.Errorf("received empty data")
	}

	log.Printf("Received complete data, size: %d bytes", buffer.Len())

	var data struct {
		Matrices     [][][]float64 `json:"matrices"`
		Coefficients []float64     `json:"coefficients"`
	}

	if err := json.Unmarshal(buffer.Bytes(), &data); err != nil {
		log.Printf("Failed to unmarshal data: %v", err)
		log.Printf("Received data: %s", buffer.String())
		return err
	}

	if len(data.Matrices) == 0 || len(data.Coefficients) == 0 {
		return fmt.Errorf("invalid data: matrices or coefficients are empty")
	}

	log.Printf("Successfully unmarshaled data: matrices=%d, coefficients=%d",
		len(data.Matrices), len(data.Coefficients))

	key, calculationTime, err := linear.ParallelLinearFormCalculation(data.Matrices, data.Coefficients, s.Wp, s.PgRepository)
	if err != nil {
		log.Printf("Error calculating linear form: %v", err)
		return err
	}

	response := &desc.GetParallelLinearFormCalculationResponse{
		Operation: "Parallel Linear Form Calculation",
		Key:       key,
		Time:      calculationTime,
	}

	if err := stream.Send(response); err != nil {
		log.Printf("Failed to send response: %v", err)
		return err
	}

	return nil
}
