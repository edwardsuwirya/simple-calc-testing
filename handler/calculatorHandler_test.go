package handler

import (
	"github.com/edwardsuwirya/calcTesting/model"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type serviceMock struct{}

func (s serviceMock) SetNumber(calculator model.Calculator) service.ICalculatorService {
	return s
}

func (s serviceMock) Result() float64 {
	return float64(3)
}

func TestNewCalculatorHandler(t *testing.T) {
	handlerMock := NewCalculatorHandler()
	assert.NotNil(t, handlerMock)
}

func TestCalculatorHandler_Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/calculator/addition", nil)

	q := req.URL.Query()
	q.Add("num1", "1")
	q.Add("num2", "2")
	req.URL.RawQuery = q.Encode()

	if err != nil {
		t.Fatal(err)
	}

	calcAddHandlerMock := NewCalculatorHandler()
	calcAddHandlerMock.SetService(serviceMock{})
	handler := http.HandlerFunc(calcAddHandlerMock.Handler)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, rr.Body.String(), "3.000000")
}
