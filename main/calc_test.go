package main

import (
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/edwardsuwirya/calcTesting/view"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculatorApp(t *testing.T) {
	t.Run("It should create main app", func(t *testing.T) {
		additionService := service.NewAdditionService()
		view := view.NewConsoleView()
		app, err := NewCalculatorApp("1", "2", additionService, view)
		assert.Nil(t, err)
		assert.NotNil(t, app)
	})

}

func TestNewCalculatorAppFailed(t *testing.T) {
	t.Run("It should return error when conversion is failed", func(t *testing.T) {
		additionService := service.NewAdditionService()
		view := view.NewConsoleView()
		_, err := NewCalculatorApp("a", "b", additionService, view)
		assert.NotNil(t, err)
	})
}

func TestNewCalculatorRun(t *testing.T) {
	t.Run("It should show result properly", func(t *testing.T) {
		additionService := service.NewAdditionService()
		view := view.NewConsoleView()
		app, _ := NewCalculatorApp("1", "2", additionService, view)
		app.run()

		res, err := app.appView.ShowResult()
		assert.Equal(t, float64(3), res)
		assert.Nil(t, err)
	})

}
