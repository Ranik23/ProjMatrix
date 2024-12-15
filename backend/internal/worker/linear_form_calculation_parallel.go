package worker

import (
	"ProjMatrix/internal/usecase/linear"
	desc "ProjMatrix/pkg/proto"
	"encoding/json"
	"google.golang.org/grpc"
	"io"
)

func (s *Service) GetParallelLinearFormCalculation(stream grpc.BidiStreamingServer[desc.GetParallelLinearFormCalculationRequest, desc.GetParallelLinearFormCalculationResponse]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		bytes := req.GetMatrix()
		var matrices [][][]float64
		err = json.Unmarshal(bytes, &matrices)
		if err != nil {
			response := &desc.GetParallelLinearFormCalculationResponse{
				Operation:    "Linear form Calculation",
				ResultMatrix: []byte{},
				Time:         0,
			}
			stream.Send(response)
			continue
		}

		result, time, err := linear.ParallelLinearFormCalculation(matrices, req.Coefficients, s.pool)
		if err != nil {
			response := &desc.GetParallelLinearFormCalculationResponse{
				Operation:    "Linear form Calculation",
				ResultMatrix: []byte{},
				Time:         0,
			}
			stream.Send(response)
			continue
		}

		resultByte, err := json.Marshal(result)
		if err != nil {
			response := &desc.GetParallelLinearFormCalculationResponse{
				Operation:    "Linear form Calculation",
				ResultMatrix: []byte{},
				Time:         0,
			}
			stream.Send(response)
			continue
		}

		response := &desc.GetParallelLinearFormCalculationResponse{
			Operation:     "Linear form Calculation",
			ResultMatrix:  resultByte,
			Time:          time,
			OperationType: "Parallel Linear Form Calculation",
		}
		stream.Send(response)
	}

	return nil
}
