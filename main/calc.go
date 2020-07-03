package main

import (
	"github.com/edwardsuwirya/calcTesting/model"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/edwardsuwirya/calcTesting/view"
)

type mainApp struct {
	appModel   model.Calculator
	appService service.ICalculatorService
	appView    view.ICalculatorView
}

func NewCalculatorApp(num01, num02 float64, service service.ICalculatorService, view view.ICalculatorView) *mainApp {
	m := model.Calculator{
		Num1: num01,
		Num2: num02,
	}
	service.SetNumber(m)
	view.SetService(service)
	return &mainApp{m, service, view}
}

func (app *mainApp) run() {
	app.appView.ShowResult()
}
func main() {
	additionService := service.NewAdditionService()
	view := view.NewConsoleView()
	NewCalculatorApp(1.0, 2.0, additionService, view).run()
}
