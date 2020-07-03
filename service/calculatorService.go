package service

import "github.com/edwardsuwirya/calcTesting/model"

type ICalculatorService interface {
	SetNumber(calculator model.Calculator) ICalculatorService
	Result() float64
}
