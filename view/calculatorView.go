package view

type ICalculatorView interface {
	ShowResult() (float64, error)
}
