package models

type multiEpisodeData struct {
	Seasionid int       `json:"season_id"`
	Episodes  []episode `json:"episodes"`
}

type episode struct {
	Aid   int    `json:"aid"`
	Cid   int    `json:"cid"`
	Title string `json:"title"`
	BVid  string `json:"bvid"`
}

type multiPageVideoData struct {
	Title string           `json:"title"`
	Pages []videoPagesData `json:"pages"`
}

type videoPagesData struct {
	Cid  int    `json:"cid"`
	Part string `json:"part"`
	Page int    `json:"page"`
}

type MultiPage struct {
	Aid       int                `json:"aid"`
	BVid      string             `json:"bvid"`
	Sections  []multiEpisodeData `json:"sections"`
	VideoData multiPageVideoData `json:"videoData"`
}

type BangumiEpData struct {
	Aid         int    `json:"aid"`
	Cid         int    `json:"cid"`
	BVid        string `json:"bvid"`
	ID          int    `json:"id"`
	EpID        int    `json:"ep_id"`
	TitleFormat string `json:"titleFormat"`
	LongTitle   string `json:"long_title"`
}

type BangumiData struct {
	EpInfo BangumiEpData   `json:"epInfo"`
	EpList []BangumiEpData `json:"epList"`
}

type BilibiliOptions struct {
	Url      string
	Html     string
	Bangumi  bool
	Aid      int
	Cid      int
	Bvid     string
	Page     int
	Subtitle string
}
