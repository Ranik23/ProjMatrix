package request

import (
	handLin "ProjMatrix/internal/api/handlers/linear"
	handPol "ProjMatrix/internal/api/handlers/polynomial"
	"ProjMatrix/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ProcessRequest(c *gin.Context, rawData map[string]interface{}) error {
	operationType, ok := rawData["operationType"].(string)
	if !ok {
		return fmt.Errorf("invalid or missing operationType")
	}

	switch operationType {
	case "manual-polynomial", "generate-polynomial":
		var polynomial entity.Polynomial
		if err := mapToStruct(rawData, &polynomial); err != nil {
			return fmt.Errorf("invalid data for a Polynomial: %w", err)
		}
		handPol.HandlePolynomial(c, &polynomial, operationType)

	case "manual-linear-form", "generate-linear-form":
		var linearForm entity.LinearForm
		if err := mapToStruct(rawData, &linearForm); err != nil {
			return fmt.Errorf("invalid data for LinearForm: %w", err)
		}
		handLin.HandleLinearForm(c, &linearForm, operationType)

	default:
		return fmt.Errorf("unknown operationType: %s", operationType)
	}

	return nil
}

func mapToStruct(data map[string]interface{}, result interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, result)
}
