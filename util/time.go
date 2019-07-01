package util

import (
	"github.com/shopspring/decimal"
	"time"
)

const (
	layout = "2006年01月" // 时间格式化格式
)

// 获取期限了所有年月格式为 xx年xx月
// startTime 起始时间  2019年10月
// num 月数
func getAllMonth(startTime string, num int) (result []string) {
	x := startTime
	t, _ := time.Parse(layout, x)
	for i := 0; i < num; i++ {
		// 获取下个月时间
		t = t.AddDate(0, 1, 0)
		// go 格式化时间2006-01-02 15:04:05 代表yyyy:MM:ss hh:mm:ss
		// 如果15改成3 会格式化成 12小时制
		result = append(result, t.Format(layout))
	}
	return result
}

func bigNumberSub(value1 float64, value2 float64) string {
	x := decimal.NewFromFloat(value1)
	y := decimal.NewFromFloat(value2)
	return x.Sub(y).String()
}
