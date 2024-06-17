package bili

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"test/models"
	"test/utils"
	"test/video/parser"
	"test/video/request"
	"test/video/videer"
)

const (
	DataTypeVideo      = "video"
	referer            = "https://www.bilibili.com"
	bilibiliBangumiAPI = "https://api.bilibili.com/pgc/player/web/playurl?"
)

var qualityString = map[int]string{
	127: "超高清 8K",
	120: "超清 4K",
	116: "高清 1080P60",
	74:  "高清 720P60",
	112: "高清 1080P+",
	80:  "高清 1080P",
	64:  "高清 720P",
	48:  "高清 720P",
	32:  "清晰 480P",
	16:  "流畅 360P",
	15:  "流畅 360P",
}

func getSubTitleCaptionPart(aid, cid int) *videer.CaptionPart {
	return nil
}

func getTitle(html string, opts models.BilibiliOptions, bilOptions videer.Options) (string, error) {
	doc, err := parser.GetDoc(html)
	if err != nil {
		return "", err
	}

	title := parser.Title(doc)
	if opts.Subtitle != "" {
		pageString := ""
		if opts.Page > 0 {
			pageString = fmt.Sprintf("P%d ", opts.Page)
		}
		if bilOptions.EpisodeTitleOnly {
			title = fmt.Sprintf("%s%s", pageString, opts.Subtitle)
		} else {
			title = fmt.Sprintf("%s %s%s", title, pageString, opts.Subtitle)
		}
	}
	return title, nil
}

func genBliApi(cid, quality int, bvid string) string {
	params := fmt.Sprintf(
		"cid=%d&bvid=%s&qn=%d&type=&otype=json&fourk=1&fnver=0&fnval=16",
		cid, bvid, quality,
	)
	return bilibiliBangumiAPI + params
}
func getStreams(options models.BilibiliOptions) (map[string]*videer.Stream, error) {
	api := genBliApi(options.Cid, 127, options.Bvid)

	jsonString, err := request.Get(api, referer, nil)
	if err != nil {
		return nil, err
	}
	var data models.Dash
	err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, err
	}

	var dashData models.DashInfo
	if data.Data.Description == nil {
		dashData = data.Result
	} else {
		dashData = data.Data
	}

	streams := make(map[string]*videer.Stream, len(dashData.Quality))

	for _, durl := range dashData.DURLs {
		var ext string
		switch dashData.DURLFormat {
		case "flv", "flv480":
			ext = "flv"
		case "mp4", "hdmp4": // nolint
			ext = "mp4"
		}

		parts := make([]*videer.Part, 0, 1)
		parts = append(parts, &videer.Part{
			URL:  durl.URL,
			Size: durl.Size,
			Ext:  ext,
		})

		streams[strconv.Itoa(dashData.CurQuality)] = &videer.Stream{
			Ext:     ext,
			Parts:   parts,
			Size:    durl.Size,
			ID:      strconv.Itoa(dashData.CurQuality),
			Quality: qualityString[dashData.CurQuality],
		}
	}
	return streams, nil
}

func bilibiliDownload(opt models.BilibiliOptions, options videer.Options) (data *videer.Data) {
	title, err := getTitle(opt.Html, opt, options)
	if err != nil {
		return videer.EmptyData(opt.Url, err)
	}
	streams, err := getStreams(opt)
	if err != nil {
		return videer.EmptyData(opt.Url, err)
	}

	data = &videer.Data{
		URL:     opt.Url,
		Site:    videer.SiteNameBilBil,
		Type:    videer.DataTypeVideo,
		Title:   title,
		Streams: streams,
		Captions: map[string]*videer.CaptionPart{
			"danmaku": {
				Part: videer.Part{
					URL: fmt.Sprintf("https://comment.bilibili.com/%d.xml", opt.Cid),
					Ext: "xml",
				},
			},
			"subtitle": getSubTitleCaptionPart(opt.Aid, opt.Cid),
		},
		Err: nil,
	}
	return data
}
func extractBangumi(url, html string, option videer.Options) ([]*videer.Data, error) {
	dataString := utils.MatchOneOf(html, `<script\s+id="__NEXT_DATA__"\s+type="application/json"\s*>(.*?)</script\s*>`)[1]
	epArrayString := utils.MatchOneOf(dataString, `"episodes"\s*:\s*(.+?)\s*,\s*"user_status"`)[1]
	fullVideoIdString := utils.MatchOneOf(dataString, `"videoId"\s*:\s*"(ep|ss)(\d+)"`)
	epSsString := fullVideoIdString[1] // "ep" or "ss"
	videoIdString := fullVideoIdString[2]

	var epArray []json.RawMessage
	err := json.Unmarshal([]byte(epArrayString), &epArray)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var data models.BangumiData
	for _, jsonByte := range epArray {
		var epInfo models.BangumiEpData
		err := json.Unmarshal(jsonByte, &epInfo)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		videoId, err := strconv.ParseInt(videoIdString, 10, 0)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// 目标集
		if epInfo.ID == int(videoId) || (epSsString == "ss" && epInfo.TitleFormat == "第1话") {
			data.EpInfo = epInfo
		}
		// 全部集
		data.EpList = append(data.EpList, epInfo)
	}
	if !option.Playlist {
		aid := data.EpInfo.Aid
		cid := data.EpInfo.Cid
		bvid := data.EpInfo.BVid
		titleFormat := data.EpInfo.TitleFormat
		longTitle := data.EpInfo.LongTitle
		if aid <= 0 || cid <= 0 || bvid == "" {
			aid = data.EpList[0].Aid
			cid = data.EpList[0].Cid
			bvid = data.EpList[0].BVid
			titleFormat = data.EpList[0].TitleFormat
			longTitle = data.EpList[0].LongTitle
		}
		options := models.BilibiliOptions{
			Url:      url,
			Html:     html,
			Bangumi:  true,
			Aid:      aid,
			Cid:      cid,
			Bvid:     bvid,
			Subtitle: fmt.Sprintf("%s %s", titleFormat, longTitle),
		}
		return []*videer.Data{bilibiliDownload(options, option)}, nil
	}
	return nil, nil
}

func extractNormalVideo(url, html string, option videer.Options) ([]*videer.Data, error) {
	return nil, nil
}

func Action(url, html string, option videer.Options) (videos []*videer.Data, err error) {
	option.ThreadNumber = 1
	if strings.Contains(url, "bangumi") {
		videos, err = extractBangumi(url, html, option)
		if err != nil {
			panic(err)
		}
	} else {
		videos, err = extractNormalVideo(url, html, option)
		if err != nil {
			panic(err)
		}
	}
	return videos, nil
}
