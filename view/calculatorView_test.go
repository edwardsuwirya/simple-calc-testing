package view

import (
	"github.com/edwardsuwirya/calcTesting/model"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

type serviceMock struct{}

func (s serviceMock) Result() float64 {
	return 10
}
func (s serviceMock) SetNumber(calculator model.Calculator) service.ICalculatorService {
	return s
}

func TestNewConsoleView(t *testing.T) {
	t.Run("It should return Console View", func(t *testing.T) {
		consoleView := NewConsoleView()
		consoleView.SetService(serviceMock{})

		assert.NotNil(t, consoleView)
	})
}

func TestConsoleView_ShowResult(t *testing.T) {
	t.Run("It should show the result", func(t *testing.T) {
		consoleView := NewConsoleView()
		consoleView.SetService(serviceMock{})
		res, err := consoleView.ShowResult()

		assert.Equal(t, float64(10), res)
		assert.Nil(t, err)
	})
}
