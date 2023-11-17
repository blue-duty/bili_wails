package bilibiliapi

import (
	"fmt"
	"testing"
)

func TestGetQrCodeUrl(t *testing.T) {
	test_url := "https://passport.bilibili.com/qrcode/getLoginUrl"
	fmt.Println(GetQrCodeUrl(test_url))
}

func TestIsLogin(t *testing.T) {
	fmt.Println(IsLogin("DedeUserID=314587828; DedeUserID__ckMd5=946acafb00d334a1; Expires=15551000; SESSDATA=ca6b146e%2C1715594401%2C817d6%2Ab1; bili_jct=e57ac55c7ebb8dd4736a90f8bcb01f2a; LIVE_BUVID=AUTO4417000424013345"))
}
