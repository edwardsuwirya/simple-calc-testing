package util

import (
	"errors"
	"github.com/edwardsuwirya/calcTesting/model"
	"strconv"
)

func CalcNumber(num01, num02 string) (model.Calculator, error) {
	num1, err1 := strconv.ParseFloat(num01, 64)
	num2, err2 := strconv.ParseFloat(num02, 64)
	if err1 != nil || err2 != nil {
		return model.Calculator{}, errors.New("Failed Conversion")
	}
	m := model.Calculator{
		Num1: num1,
		Num2: num2,
	}
	return m, nil
}
