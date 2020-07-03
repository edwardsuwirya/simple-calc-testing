package route

import (
	"github.com/edwardsuwirya/calcTesting/handler"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/gorilla/mux"
	"net/http"
)

type ICalculatorRoute interface {
	SetRouter(r *mux.Router) ICalculatorRoute
	Init()
}

type calculatorRoute struct {
	route *mux.Router
}

func (cr *calculatorRoute) SetRouter(r *mux.Router) ICalculatorRoute {
	cr.route = r
	return cr
}
func (cr *calculatorRoute) Init() {
	calc := cr.route.PathPrefix("/calculator").Subrouter()
	additionService := service.NewAdditionService()
	calc.HandleFunc("/addition", handler.NewCalculatorHandler().SetService(additionService).Handler).Methods(http.MethodGet)
	calc.HandleFunc("/substraction", handler.NewCalculatorHandler().SetService(additionService).Handler).Methods(http.MethodGet)
	calc.HandleFunc("/multiplication", handler.NewCalculatorHandler().SetService(additionService).Handler).Methods(http.MethodGet)
	calc.HandleFunc("/division", handler.NewCalculatorHandler().SetService(additionService).Handler).Methods(http.MethodGet)
}

func NewCalculatorRoute() ICalculatorRoute {
	return &calculatorRoute{}
}
