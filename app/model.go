package app

type VideoInfo struct {
	Bvid      string `json:"bvid"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Pic       string `json:"pic"`
	Pubdate   string `json:"pubdate"`
	OwnerName string `json:"owner_name"`
	OwnerFace string `json:"owner_face"`

	//  基本信息 播放 弹幕 点赞 投币 收藏 分享 评论
	Stat Stat `json:"stat"`
	//  视频分P信息
	Eps []Episode `json:"eps"`
}

type Stat struct {
	View     string `json:"view"`
	Danmaku  string `json:"danmaku"`
	Like     string `json:"like"`
	Coin     string `json:"coin"`
	Favorite string `json:"favorite"`
	Share    string `json:"share"`
	Reply    string `json:"reply"`
}

type Sections struct {
	SeasonID int       `json:"season_id"`
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Episodes []Episode `json:"episodes"`
}

type Episode struct {
	Aid      int    `json:"aid"`
	Bvid     string `json:"bvid"`
	Cid      int    `json:"cid"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Index    int    `json:"index"`
	Selected bool   `json:"selected"`
}
