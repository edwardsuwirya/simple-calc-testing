package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type serviceMock struct{}

func (serviceMock) Result() float64 {
	return 10
}

func TestNewConsoleView(t *testing.T) {
	t.Run("It should return Console View", func(t *testing.T) {
		consoleView := NewConsoleView(serviceMock{})

		assert.NotNil(t, consoleView)
	})
}

func TestConsoleView_ShowResult(t *testing.T) {
	t.Run("It should show the result", func(t *testing.T) {
		consoleView := NewConsoleView(serviceMock{})
		res, err := consoleView.ShowResult()

		assert.Equal(t, 10, res)
		assert.NotNil(t, err)
	})
}
