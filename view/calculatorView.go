package view

import "github.com/edwardsuwirya/calcTesting/service"

type ICalculatorView interface {
	SetService(service service.ICalculatorService) ICalculatorView
	ShowResult() (float64, error)
}
