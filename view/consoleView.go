package view

import (
	"github.com/edwardsuwirya/calcTesting/service"
	"log"
)

type consoleView struct {
	calculatorService service.ICalculatorService
}

func NewConsoleView() ICalculatorView {
	return &consoleView{}
}

func (cv *consoleView) SetService(service service.ICalculatorService) ICalculatorView {
	cv.calculatorService = service
	return cv
}
func (cv *consoleView) ShowResult() (float64, error) {
	res := cv.calculatorService.Result()
	log.Print(res)
	return res, nil
}
