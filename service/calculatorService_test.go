package service

import (
	"github.com/edwardsuwirya/calcTesting/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAdditionService(t *testing.T) {
	t.Run("It should return Calculator service", func(t *testing.T) {
		modelMock := model.Calculator{}

		calculatorService := NewAdditionService()
		calculatorService.SetNumber(modelMock)

		assert.NotNil(t, calculatorService)
	})
}

type modelCalculatorMock struct {
	calculator     model.Calculator
	expectedResult float64
}

func TestAdditionService_Result(t *testing.T) {
	//Test Table
	tests := []modelCalculatorMock{
		{
			calculator: model.Calculator{
				Num1: 1,
				Num2: 2,
			},
			expectedResult: 3,
		},
		{
			calculator: model.Calculator{
				Num1: 1,
				Num2: -2,
			},
			expectedResult: -1,
		},
		{
			calculator: model.Calculator{
				Num1: -4,
				Num2: 2,
			},
			expectedResult: -2,
		},

		{
			calculator: model.Calculator{
				Num1: -7,
				Num2: -3,
			},
			expectedResult: -10,
		},
	}
	for _, test := range tests {
		t.Run("It should return addition result", func(t *testing.T) {
			calculatorService := NewAdditionService()
			calculatorService.SetNumber(test.calculator)

			expected := test.expectedResult
			actual := calculatorService.Result()

			assert.Equal(t, actual, expected)
		})
	}
}

func TestAdditionServiceWithoutNumber_Result(t *testing.T) {
	t.Run("It should return 0 when no nums provided", func(t *testing.T) {
		modelMock := model.Calculator{}

		calculatorService := NewAdditionService()
		calculatorService.SetNumber(modelMock)

		expected := float64(0)
		actual := calculatorService.Result()

		assert.Equal(t, actual, expected)
	})
}
