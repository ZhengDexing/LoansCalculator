package util

import (
	"time"
)

// 获取期限了所有年月格式为 xx年xx月
// num 月数
func getAllMonth (num int) (result []string) {
	t := time.Now()
	for i:=0;i<num;i++{
		// 获取下个月时间
		t = t.AddDate(0, 1, 0)
		// go 格式化时间2006-01-02 15:04:05 代表yyyy:MM:ss hh:mm:ss
		// 如果15改成3 会格式化成 12小时制
		result = append(result, t.Format("2006年01月"))
	}
	return result
}
