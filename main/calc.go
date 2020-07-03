package main

import (
	"fmt"
	"github.com/edwardsuwirya/calcTesting/model"
	myRoute "github.com/edwardsuwirya/calcTesting/route"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/edwardsuwirya/calcTesting/util"
	"github.com/edwardsuwirya/calcTesting/view"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type mainApp struct {
	appModel   model.Calculator
	appService service.ICalculatorService
	appRoute   myRoute.ICalculatorRoute
	appView    view.ICalculatorView
}

func NewCalculatorApp(num01, num02 string, service service.ICalculatorService, view view.ICalculatorView) (*mainApp, error) {
	m, _ := util.CalcNumber(num01, num02)
	service.SetNumber(m)
	view.SetService(service)
	return &mainApp{
		appModel:   m,
		appService: service,
		appRoute:   nil,
		appView:    view,
	}, nil
}

func NewCalculatorAppServer(route myRoute.ICalculatorRoute) *mainApp {
	return &mainApp{
		appModel:   model.Calculator{},
		appService: nil,
		appRoute:   route,
		appView:    nil,
	}
}

func (app *mainApp) run() {
	app.appView.ShowResult()
}

func (app *mainApp) runServer(host, port string) {
	r := mux.NewRouter()

	app.appRoute.SetRouter(r).Init()

	log.Print("Server is listening")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r); err != nil {
		log.Panic(err)
	}
}
func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Can not run")
		}
	}()

	if len(os.Args) == 3 {
		additionService := service.NewAdditionService()
		view := view.NewConsoleView()
		appConsole, err := NewCalculatorApp(os.Args[1], os.Args[2], additionService, view)
		if err != nil {
			panic("Failed to run")
		}
		appConsole.run()
	} else {
		route := myRoute.NewCalculatorRoute()
		appServer := NewCalculatorAppServer(route)
		appServer.runServer("localhost", "7000")
	}

}
