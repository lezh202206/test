package videer

const (
	VideoTypeBangumi = "bangumi"

	SiteNameBilBil = "哔哩哔哩 bilibili.com"

	DataTypeVideo = "video"
)

type DataType string

type Part struct {
	URL  string `json:"url"`
	Size int64  `json:"size"`
	Ext  string `json:"ext"`
}

type Stream struct {
	// eg: "1080"
	ID string `json:"id"`
	// eg: "1080P xxx"
	Quality string `json:"quality"`
	// [Part: {URL, Size, Ext}, ...]
	// Some video stream have multiple parts,
	// and can also be used to download multiple image files at once
	Parts []*Part `json:"parts"`
	// total size of all urls
	Size int64 `json:"size"`
	// the file extension after video parts merged
	Ext string `json:"ext"`
	// if the parts need mux
	NeedMux bool
}

type CaptionPart struct {
	Part
	Transform func([]byte) ([]byte, error) `json:"-"`
}

type Data struct {
	// URL is used to record the address of this download
	URL   string   `json:"url"`
	Site  string   `json:"site"`
	Title string   `json:"title"`
	Type  DataType `json:"type"`
	// each stream has it's own Parts and Quality
	Streams map[string]*Stream `json:"streams"`
	// danmaku, subtitles, etc
	Captions map[string]*CaptionPart `json:"caption"`
	// Err is used to record whether an error occurred when extracting the list data
	Err error `json:"err"`
}

type Options struct {
	// Playlist indicates if we need to extract the whole playlist rather than the single video.
	Playlist bool
	// Items defines wanted items from a playlist. Separated by commas like: 1,5,6,8-10.
	Items string
	// ItemStart defines the starting item of a playlist.
	ItemStart int
	// ItemEnd defines the ending item of a playlist.
	ItemEnd int

	// ThreadNumber defines how many threads will use in the extraction, only works when Playlist is true.
	ThreadNumber int
	Cookie       string

	// EpisodeTitleOnly indicates file name of each bilibili episode doesn't include the playlist title
	EpisodeTitleOnly bool

	//YoukuCcode    string
	//YoukuCkey     string
	//YoukuPassword string
}

// EmptyData returns an "empty" Data object with the given URL and error.
func EmptyData(url string, err error) *Data {
	return &Data{
		URL: url,
		Err: err,
	}
}
