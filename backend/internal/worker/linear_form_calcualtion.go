package worker

import (
	"ProjMatrix/internal/usecase/linear"
	desc "ProjMatrix/pkg/proto"
	"context"
	"github.com/goccy/go-json"
)

func (s *Service) GetLinearFormCalculation(ctx context.Context, req *desc.GetLinearFormCalculationRequest) (*desc.GetLinearFormCalculationResponse, error) {
	data := req.GetMatrix()
	var matrices [][][]float64
	err := json.Unmarshal(data, &matrices)
	if err != nil {
		return &desc.GetLinearFormCalculationResponse{
			Operation:    "Linear form Calculation",
			ResultMatrix: []byte{},
			Time:         0,
		}, err
	}

	if req.Coefficients == nil {
		req.Coefficients = []float64{}
	}

	result, time, err := linear.LinearFormCalculation(matrices, req.Coefficients)
	if err != nil {
		return &desc.GetLinearFormCalculationResponse{
			Operation:    "Linear form Calculation",
			ResultMatrix: []byte{},
			Time:         0,
		}, err
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		return &desc.GetLinearFormCalculationResponse{
			Operation:    "Linear form Calculation",
			ResultMatrix: []byte{},
			Time:         0,
		}, err
	}

	return &desc.GetLinearFormCalculationResponse{
		Operation:     "Linear form Calculation",
		ResultMatrix:  resultByte,
		Time:          time,
		OperationType: "Linear Form Calculation",
	}, nil
}
