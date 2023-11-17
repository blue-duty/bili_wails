package bilibiliapi

type VideoInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    Data   `json:"data"`
}

type VideoInfoForView struct {
	Bvid      string    `json:"bvid"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Pic       string    `json:"pic"`
	Owner     Owner     `json:"owner"`
	UgcSeason UgcSeason `json:"ugc_season"`
}

type Data struct {
	Bvid               string      `json:"bvid"`
	Aid                int         `json:"aid"`
	Videos             int         `json:"videos"`
	Tid                int         `json:"tid"`
	Tname              string      `json:"tname"`
	Copyright          int         `json:"copyright"`
	Pic                string      `json:"pic"`
	Title              string      `json:"title"`
	Pubdate            int64       `json:"pubdate"`
	Ctime              int64       `json:"ctime"`
	Desc               string      `json:"desc"`
	DescV2             []DescV2    `json:"desc_v2"`
	State              int         `json:"state"`
	Duration           int         `json:"duration"`
	Rights             Rights      `json:"rights"`
	Owner              Owner       `json:"owner"`
	Stat               Stat        `json:"stat"`
	Dynamic            string      `json:"dynamic"`
	Cid                int         `json:"cid"`
	Dimension          Dimension   `json:"dimension"`
	SeasonID           int         `json:"season_id"`
	Premiere           interface{} `json:"premiere"`
	TeenageMode        int         `json:"teenage_mode"`
	IsChargeableSeason bool        `json:"is_chargeable_season"`
	IsStory            bool        `json:"is_story"`
	IsUpowerExclusive  bool        `json:"is_upower_exclusive"`
	IsUpowerPlay       bool        `json:"is_upower_play"`
	EnableVt           int         `json:"enable_vt"`
	VtDisplay          string      `json:"vt_display"`
	NoCache            bool        `json:"no_cache"`
	Pages              []Page      `json:"pages"`
	Subtitle           interface{} `json:"subtitle"`
	UgcSeason          UgcSeason   `json:"ugc_season"`
	IsSeasonDisplay    bool        `json:"is_season_display"`
	UserGarb           UserGarb    `json:"user_garb"`
	HonorReply         interface{} `json:"honor_reply"`
	LikeIcon           string      `json:"like_icon"`
	NeedJumpBv         bool        `json:"need_jump_bv"`
	DisableShowUpInfo  bool        `json:"disable_show_up_info"`
}

type DescV2 struct {
	RawText string `json:"raw_text"`
	Type    int    `json:"type"`
	BizID   int    `json:"biz_id"`
}

type Rights struct {
	Bp            int `json:"bp"`
	Elec          int `json:"elec"`
	Download      int `json:"download"`
	Movie         int `json:"movie"`
	Pay           int `json:"pay"`
	Hd5           int `json:"hd5"`
	NoReprint     int `json:"no_reprint"`
	Autoplay      int `json:"autoplay"`
	UgcPay        int `json:"ugc_pay"`
	IsCooperation int `json:"is_cooperation"`
	UgcPayPreview int `json:"ugc_pay_preview"`
	ArcPay        int `json:"arc_pay"`
	FreeWatch     int `json:"free_watch"`
	CleanMode     int `json:"clean_mode"`
	IsSteinGate   int `json:"is_stein_gate"`
	Is360         int `json:"is_360"`
	NoShare       int `json:"no_share"`
	NoBackground  int `json:"no_background"`
}

type Owner struct {
	Mid  int    `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}

type Stat struct {
	Aid        int    `json:"aid"`
	View       int    `json:"view"`
	Danmaku    int    `json:"danmaku"`
	Reply      int    `json:"reply"`
	Favorite   int    `json:"favorite"`
	Coin       int    `json:"coin"`
	Share      int    `json:"share"`
	NowRank    int    `json:"now_rank"`
	HisRank    int    `json:"his_rank"`
	Like       int    `json:"like"`
	Dislike    int    `json:"dislike"`
	Evaluation string `json:"evaluation"`
	ArgueMsg   string `json:"argue_msg"`
	Vt         int    `json:"vt"`
}

type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Rotate int `json:"rotate"`
}

type Page struct {
	Cid        int       `json:"cid"`
	Page       int       `json:"page"`
	From       string    `json:"from"`
	Part       string    `json:"part"`
	Duration   int       `json:"duration"`
	Vid        string    `json:"vid"`
	Weblink    string    `json:"weblink"`
	Dimension  Dimension `json:"dimension"`
	FirstFrame string    `json:"first_frame"`
}

type UgcSeason struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Mid       int       `json:"mid"`
	Intro     string    `json:"intro"`
	SignState int       `json:"sign_state"`
	Attribute int       `json:"attribute"`
	Sections  []Section `json:"sections"`
}

type Section struct {
	SeasonID int       `json:"season_id"`
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Type     int       `json:"type"`
	Episodes []Episode `json:"episodes"`
}

type Episode struct {
	SeasonID  int    `json:"season_id"`
	SectionID int    `json:"section_id"`
	ID        int    `json:"id"`
	Aid       int    `json:"aid"`
	Cid       int    `json:"cid"`
	Title     string `json:"title"`
	Attribute int    `json:"attribute"`
	Arc       Arc    `json:"arc"`
	Page      Page   `json:"page"`
	Bvid      string `json:"bvid"`
}

type Arc struct {
	Aid       int         `json:"aid"`
	Videos    int         `json:"videos"`
	TypeID    int         `json:"type_id"`
	TypeName  string      `json:"type_name"`
	Copyright int         `json:"copyright"`
	Pic       string      `json:"pic"`
	Title     string      `json:"title"`
	Pubdate   int64       `json:"pubdate"`
	Ctime     int64       `json:"ctime"`
	Desc      string      `json:"desc"`
	State     int         `json:"state"`
	Duration  int         `json:"duration"`
	Rights    Rights      `json:"rights"`
	Author    Author      `json:"author"`
	Stat      Stat        `json:"stat"`
	Dynamic   string      `json:"dynamic"`
	Dimension Dimension   `json:"dimension"`
	DescV2    interface{} `json:"desc_v2"`
}

type Author struct {
	Mid  int    `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}

type UserGarb struct {
	UrlImageAniCut string `json:"url_image_ani_cut"`
}
