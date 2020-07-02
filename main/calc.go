package main

import (
	"github.com/edwardsuwirya/calcTesting/model"
	"github.com/edwardsuwirya/calcTesting/service"
	"github.com/edwardsuwirya/calcTesting/view"
)

func main() {
	c := model.Calculator{
		Num1: 1,
		Num2: 2,
	}
	additionService := service.NewAdditionService(c)
	view.NewConsoleView(additionService).ShowResult()
}
