package view

import "testing"

type serviceMock struct{}

func (serviceMock) Result() float64 {
	return 10
}

func TestNewConsoleView(t *testing.T) {
	t.Run("It should return Console View", func(t *testing.T) {
		consoleView := NewConsoleView(serviceMock{})

		if consoleView == nil {
			t.Error("Can not be nil")
		}
	})
}

func TestConsoleView_ShowResult(t *testing.T) {
	t.Run("It should show the result", func(t *testing.T) {
		consoleView := NewConsoleView(serviceMock{})
		res, err := consoleView.ShowResult()

		if res != 10 {
			t.Error("Result is wrong")
		}

		if err != nil {
			t.Error("error is occured")
		}
	})
}
