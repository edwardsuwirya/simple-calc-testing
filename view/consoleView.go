package view

import (
	"github.com/edwardsuwirya/calcTesting/service"
	"log"
)

type consoleView struct {
	calculatorService service.ICalculatorService
}

func NewConsoleView(service service.ICalculatorService) ICalculatorView {
	return &consoleView{calculatorService: service}
}

func (cv *consoleView) ShowResult() (float64, error) {
	res := cv.calculatorService.Result()
	log.Print(res)
	return res, nil
}
