package bilibiliapi

import (
	bilibiliapi "bili/backend/bilibili_api"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 获取vedio信息
func GetVideoInfo(bvid string) (VideoInfo, error) {
	url := "https://api.bilibili.com/x/web-interface/view?bvid=%s"
	url = fmt.Sprintf(url, bvid)
	var videoInfo VideoInfo
	fmt.Println(url)
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	// req.Header.Set("User-Agent", UserAgent)
	// set referer
	req.Header.Add("Origin", "https://www.bilibili.com")
	req.Header.Add("Referer", "https://www.bilibili.com")
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", bilibiliapi.UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting video info:", err)
		return VideoInfo{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return VideoInfo{}, err
	}
	err = json.Unmarshal(body, &videoInfo)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return VideoInfo{}, err
	}
	if videoInfo.Code != 0 {
		return VideoInfo{}, fmt.Errorf("error getting video info: %s", videoInfo.Message)
	}
	return videoInfo, nil
}
