package util

import "github.com/shopspring/decimal"

// 精度运算
func BigNumberSub(value1 float64, value2 float64) string {
	x := decimal.NewFromFloat(value1)
	y := decimal.NewFromFloat(value2)
	return x.Sub(y).String()
}
