package main

import (
	"bili/app"
	bilibiliapi "bili/backend/bilibili_api"
	video "bili/backend/bilibili_api/video"
	"bili/backend/engine"
	"bili/backend/parser"
	"bili/backend/persist"
	"bili/backend/scheduler"
	"bili/backend/tool"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"

	goruntime "runtime"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
// NewApp 创建一个新的 App 应用程序
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) DomReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
// func (a *App) shutdown(ctx context.Context) {
// 	// Perform your teardown here
// 	// 在此处做一些资源释放的操作

// }

// 程序退出
func (a *App) Quit() {
	runtime.Quit(a.ctx)
}

// 最小化窗口
func (a *App) Minimize() {
	runtime.WindowMinimise(a.ctx)
}

// 打开目录选择器
func (a *App) OpenDir() string {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择目录",
	})
	if err != nil {
		return ""
	}
	return dir
}

// 获取剪贴板内容
func (a *App) GetClipboard() string {
	clipboard, err := runtime.ClipboardGetText(a.ctx)
	if err != nil {
		return ""
	}
	fmt.Println(clipboard)
	return clipboard
}

// 下载视频
func (a *App) DownloadVedio(url string, dir string) string {
	itemProcessFun := persist.GetItemProcessFun()
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	itemChan, err := itemProcessFun(&wg)
	if err != nil {
		return fmt.Sprintf("下载失败: %s", err)
	}
	var req *engine.Request
	// 获取bv号
	bvid := tool.GetBvid(url)
	aid := parser.Bv2av(bvid)
	req = parser.GetRequestByAid(aid, dir)
	queueScheduler := scheduler.NewConcurrentScheduler()
	conEngine := engine.NewConcurrentEngine(30, queueScheduler, itemChan)

	conEngine.Run(req)
	wg.Wait()
	return "下载完成"
}

// 打开文件管理器当前位置: Windows
func (a *App) OpenFileExplorer(dir string) {
	var cmd *exec.Cmd
	cos := goruntime.GOOS
	if cos == "darwin" {
		cmd = exec.Command("open", dir)
	} else if cos == "windows" {
		//  将目录转为Windows格式
		dir = strings.Replace(dir, "/", "\\", -1)
		cmd = exec.Command("explorer", dir)
	} else {
		cmd = exec.Command("xdg-open", dir)
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

var AccessKey string

// bilibili扫码登录
func (a *App) BiliLogin() string {
	// key, url := bilibiliapi.GetLoginKeyAndLoginUrl()
	// AccessKey = key
	// return bilibiliapi.GetQrCodeUrl(url)
	return "https://www.bilibili.com/"
}

// 获取cookie
func (a *App) GetCookie() {
	bilibiliapi.VerifyLogin(AccessKey, "./cookie.txt")
}

// 获取头像数据
func (a *App) GetAvatar() string {
	data := bilibiliapi.GetImage("https://i1.hdslb.com/bfs/face/8cb23c7e052799e94cfebd73035a316a26979bea.jpg@120w_120h_1c.webp")
	// base64Png := base64.StdEncoding.EncodeToString(data)
	// return "data:image/png;base64," + base64Png
	return data
}

func (a *App) GetVideoInfo(url string) app.VideoInfo {
	// 获取bv号
	bvid := tool.GetBvid(url)
	videoInfo, err := video.GetVideoInfo(bvid)
	if err != nil {
		fmt.Println(err)
	}
	// 获取头像
	avatar := bilibiliapi.GetImage(videoInfo.Data.Owner.Face)
	// 获取视频封面
	cover := bilibiliapi.GetImage(videoInfo.Data.Pic)
	// 投稿时间 -> 转为年月日 时分秒
	t := time.Unix(videoInfo.Data.Pubdate, 0).Format("2006-01-02 15:04:05")

	//  获取视频分P信息
	var eps []app.Episode
	var index int
	for _, section := range videoInfo.Data.UgcSeason.Sections {
		for _, episode := range section.Episodes {
			index++
			eps = append(eps, app.Episode{
				Aid:      episode.Aid,
				Bvid:     episode.Bvid,
				Cid:      episode.Cid,
				Title:    episode.Title,
				Duration: tool.GetDuration(episode.Arc.Duration),
				Index:    index,
			})
		}
	}

	return app.VideoInfo{
		Bvid:      videoInfo.Data.Bvid,
		Title:     videoInfo.Data.Title,
		Desc:      videoInfo.Data.Desc,
		OwnerName: videoInfo.Data.Owner.Name,
		Pic:       cover,
		OwnerFace: avatar,
		Pubdate:   t,
		Stat: app.Stat{
			Danmaku:  tool.IsOverTenThousand(videoInfo.Data.Stat.Danmaku),
			View:     tool.IsOverTenThousand(videoInfo.Data.Stat.View),
			Like:     tool.IsOverTenThousand(videoInfo.Data.Stat.Like),
			Coin:     tool.IsOverTenThousand(videoInfo.Data.Stat.Coin),
			Favorite: tool.IsOverTenThousand(videoInfo.Data.Stat.Favorite),
			Share:    tool.IsOverTenThousand(videoInfo.Data.Stat.Share),
			Reply:    tool.IsOverTenThousand(videoInfo.Data.Stat.Reply),
		},
		Eps: eps,
	}
}
