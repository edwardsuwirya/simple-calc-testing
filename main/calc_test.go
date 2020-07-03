package main

import (
	"github.com/edwardsuwirya/calcTesting/model"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/edwardsuwirya/calcTesting/view"
	"github.com/stretchr/testify/assert"
	"testing"
)

type serviceMock struct {
}

func (s serviceMock) SetNumber(calculator model.Calculator) service.ICalculatorService {
	return s
}

func (s serviceMock) Result() float64 {
	return float64(3)
}

type viewMock struct{}

func (v viewMock) SetService(service service.ICalculatorService) view.ICalculatorView {
	return v
}

func (v viewMock) ShowResult() (float64, error) {
	return float64(3), nil
}

func TestNewCalculatorApp(t *testing.T) {
	t.Run("It should create main app", func(t *testing.T) {
		app, err := NewCalculatorApp("1", "2", serviceMock{}, viewMock{})
		assert.Nil(t, err)
		assert.NotNil(t, app)
	})
}

func TestNewCalculatorAppFailed(t *testing.T) {
	t.Run("It should return error when conversion is failed", func(t *testing.T) {
		_, err := NewCalculatorApp("a", "b", serviceMock{}, viewMock{})
		assert.NotNil(t, err)
	})
}

func TestNewCalculatorRun(t *testing.T) {
	t.Run("It should show result properly", func(t *testing.T) {
		//Kasih test table
		app, _ := NewCalculatorApp("1", "2", serviceMock{}, viewMock{})
		app.run()

		res, err := app.appView.ShowResult()
		assert.Equal(t, float64(3), res)
		assert.Nil(t, err)
	})

}
