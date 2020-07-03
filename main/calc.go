package main

import (
	"errors"
	"github.com/edwardsuwirya/calcTesting/model"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/edwardsuwirya/calcTesting/view"
	"os"
	"strconv"
)

type mainApp struct {
	appModel   model.Calculator
	appService service.ICalculatorService
	appView    view.ICalculatorView
}

func NewCalculatorApp(num01, num02 string, service service.ICalculatorService, view view.ICalculatorView) (*mainApp, error) {
	num1, err1 := strconv.ParseFloat(num01, 64)
	num2, err2 := strconv.ParseFloat(num02, 64)
	if err1 != nil || err2 != nil {
		return nil, errors.New("Failed Conversion")
	}
	m := model.Calculator{
		Num1: num1,
		Num2: num2,
	}
	service.SetNumber(m)
	view.SetService(service)
	return &mainApp{m, service, view}, nil
}

func (app *mainApp) run() {
	app.appView.ShowResult()
}
func main() {
	additionService := service.NewAdditionService()
	view := view.NewConsoleView()

	app, err := NewCalculatorApp(os.Args[1], os.Args[2], additionService, view)

	//defer func() {
	//	if e := recover(); e != nil {
	//		fmt.Println("Can not run")
	//	}
	//}()
	if err != nil {
		panic("Failed to run")
	}
	app.run()
}
