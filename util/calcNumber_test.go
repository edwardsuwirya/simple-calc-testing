package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcNumber(t *testing.T) {
	t.Run("It should return calculator model", func(t *testing.T) {
		num1 := "1"
		num2 := "2"
		expectedModel, err := CalcNumber(num1, num2)
		assert.Nil(t, err)
		assert.Equal(t, expectedModel.Num1, float64(1))
	})
}
