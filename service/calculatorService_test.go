package service

import (
	"github.com/edwardsuwirya/calcTesting/model"
	"testing"
)

func TestNewAdditionService(t *testing.T) {
	t.Run("It should return Calculator service", func(t *testing.T) {
		modelMock := model.Calculator{
			Num1: 1,
			Num2: 2,
		}

		calculatorService := NewAdditionService(modelMock)

		if calculatorService == nil {
			t.Error("Can not be null")
		}
	})
}

func TestAdditionService_Result(t *testing.T) {
	t.Run("It should return addition result", func(t *testing.T) {
		modelMock := model.Calculator{
			Num1: 1,
			Num2: 2,
		}

		calculatorService := NewAdditionService(modelMock)

		expected := float64(3)
		actual := calculatorService.Result()

		if actual != expected {
			t.Error("Calculation is failed")
		}

	})
}
