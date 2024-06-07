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
	"test/video/videer"
)

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
func getStreams(aid, cid int) map[string]*videer.Stream {
	return nil
}

func bilibiliDownload(opt models.BilibiliOptions, options videer.Options) (data *videer.Data) {
	title, err := getTitle(opt.Html, opt, options)
	if err != nil {
		return videer.EmptyData(opt.Url, err)
	}

	data = &videer.Data{
		URL:   opt.Url,
		Site:  videer.SiteNameBilBil,
		Type:  videer.DataTypeVideo,
		Title: title,
		//Streams: getStreams(),
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

func Action(url, html string, option videer.Options) ([]*videer.Data, error) {
	option.ThreadNumber = 1
	if strings.Contains(url, "bangumi") {
		// handle bangumi
		return extractBangumi(url, html, option)
	} else {
		// handle normal video
		return extractNormalVideo(url, html, option)
	}
}
