package tool

import (
	"strings"
)

func GetAppKey(entropy string) (appkey, sec string) {
	revEntropy := ReverseRunes([]rune(entropy))
	for i := range revEntropy {
		revEntropy[i] = revEntropy[i] + 2
	}
	ret := strings.Split(string(revEntropy), ":")

	return ret[0], ret[1]
}

func ReverseRunes(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return runes
}

// https://www.bilibili.com/video/BV1d34y1M7ki/?spm_id_from=333.1007.tianma.1-1-1.click
// 获取视频BV号 ->	BV1d34y1M7ki
func GetBvid(url string) string {
	// 去掉问号后面的参数
	url = strings.Split(url, "?")[0]
	bvid := strings.Split(url, "/")[4]
	return bvid
}
