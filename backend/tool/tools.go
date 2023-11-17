package tool

import "fmt"

// 判断是否超过万，超过万则转换为万为单位
func IsOverTenThousand(num int) string {
	if num > 10000 {
		return fmt.Sprintf("%.1fw", float64(num)/10000)
	}
	return fmt.Sprintf("%d", num)
}

// 将时间戳转为几分几秒
func GetDuration(duration int) string {
	minute := duration / 60
	second := duration % 60
	return fmt.Sprintf("%dm%ds", minute, second)
}
