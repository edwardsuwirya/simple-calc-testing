package main

import (
	"github.com/edwardsuwirya/calcTesting/model"
	"github.com/edwardsuwirya/calcTesting/service"
	"log"
)

func main() {
	c := model.Calculator{
		Num1: 1,
		Num2: 2,
	}
	additionResult := service.NewAdditionService(c).Result()
	log.Print(additionResult)
}
