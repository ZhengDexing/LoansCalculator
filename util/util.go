package util

import (
	"math"

	"github.com/shopspring/decimal"
)

// BigNumberAdd 精度运算 加法
func BigNumberAdd(value1 float64, value2 float64) string {
	x := decimal.NewFromFloat(value1)
	y := decimal.NewFromFloat(value2)
	return x.Add(y).String()
}

// BigNumberSub 精度运算 减法
func BigNumberSub(value1 float64, value2 float64) string {
	x := decimal.NewFromFloat(value1)
	y := decimal.NewFromFloat(value2)
	return x.Sub(y).String()
}

// BigNumberMul 精度运算 乘法
func BigNumberMul(value1 float64, value2 float64) string {
	x := decimal.NewFromFloat(value1)
	y := decimal.NewFromFloat(value2)
	return x.Mul(y).String()
}

// BigNumberDiv 精度运算 除法
func BigNumberDiv(value1 float64, value2 float64) string {
	x := decimal.NewFromFloat(value1)
	y := decimal.NewFromFloat(value2)
	return x.Div(y).String()
}

// Powerf2 递归法 求x^n
func Powerf2(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else {
		return x * Powerf2(x, n-1)
	}
}

func Round(f float64) float64 {
	n10 := math.Pow10(2)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
