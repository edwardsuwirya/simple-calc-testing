package handler

import (
	"github.com/edwardsuwirya/calcTesting/service"
	"net/http"
)

type IHandler interface {
	SetService(service service.ICalculatorService) IHandler
	Handler(http.ResponseWriter, *http.Request)
}
