package service

import (
	"github.com/edwardsuwirya/calcTesting/model"
)

type additionService struct {
	calculator model.Calculator
}

func NewAdditionService() ICalculatorService {
	model := model.Calculator{}
	return &additionService{calculator: model}
}

func (as *additionService) SetNumber(calculator model.Calculator) ICalculatorService {
	as.calculator = calculator
	return as
}

func (as *additionService) Result() float64 {
	return as.calculator.Num1 + as.calculator.Num2
}
