package handler

import (
	"fmt"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/edwardsuwirya/calcTesting/util"
	"net/http"
)

type calculatorHandler struct {
	service service.ICalculatorService
}

func (c *calculatorHandler) Handler(w http.ResponseWriter, r *http.Request) {
	mapNums := r.URL.Query()
	m, _ := util.CalcNumber(mapNums["num1"][0], mapNums["num2"][0])
	result := c.service.SetNumber(m).Result()
	w.Write([]byte(fmt.Sprintf("%f", result)))
}

func (c *calculatorHandler) SetService(service service.ICalculatorService) IHandler {
	c.service = service
	return c
}
func NewCalculatorHandler() IHandler {
	return &calculatorHandler{}
}
