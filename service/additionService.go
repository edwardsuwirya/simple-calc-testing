package service

import "github.com/edwardsuwirya/calcTesting/model"

type additionService struct {
	calculator model.Calculator
}

func NewAdditionService(model model.Calculator) ICalculatorService {
	return &additionService{calculator: model}
}

func (as *additionService) Result() float64 {
	return as.calculator.Num1 + as.calculator.Num2
}
