package downloader

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/iawia002/lux/utils"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"test/video/request"
	"test/video/videer"
	"time"
)

const (
	DOWNLOAD_FILE_EXT = ".download"
)

type Options struct {
	InfoOnly       bool
	Silent         bool
	Stream         string
	AudioOnly      bool
	Refer          string
	OutputPath     string
	OutputName     string
	FileNameLength int
	Caption        bool

	MultiThread  bool
	ThreadNumber int
	RetryTimes   int
	ChunkSizeMB  int
	// Aria2
	UseAria2RPC bool
	Aria2Token  string
	Aria2Method string
	Aria2Addr   string
}

type Downloader struct {
	bar    *pb.ProgressBar
	option Options
}

func New(option Options) *Downloader {
	downloader := &Downloader{
		option: option,
	}
	return downloader
}

func progressBar(size int64) *pb.ProgressBar {
	tmpl := `{{counters .}} {{bar . "[" "=" ">" "-" "]"}} {{speed .}} {{percent . | green}} {{rtime .}}`
	return pb.New64(size).
		Set(pb.Bytes, true).
		SetMaxWidth(1000).
		SetTemplate(pb.ProgressBarTemplate(tmpl))
}

func (downloader Downloader) getStream(data *videer.Data) *videer.Stream {
	sortedStreams := genSortedStreams(data.Streams)

	streamName := downloader.option.Stream
	if streamName == "" {
		streamName = sortedStreams[0].ID
	}
	stream, ok := data.Streams[streamName]
	if !ok {
		return nil
	}
	return stream
}

func (downloader Downloader) Download(data *videer.Data) error {
	title := utils.FileName(data.Title, "", downloader.option.FileNameLength)

	stream := downloader.getStream(data)
	if stream == nil {
		return errors.Errorf("no stream named")
	}

	if !downloader.option.Silent {
		printStreamInfo(data, stream)
	}

	mergedFilePath, err := utils.FilePath(title, stream.Ext, downloader.option.FileNameLength, downloader.option.OutputPath, false)
	if err != nil {
		return err
	}
	_, mergedFileExists, err := utils.FileSize(mergedFilePath)
	if err != nil {
		return err
	}
	if mergedFileExists {
		return errors.Errorf("%s: file already exists, skipping\n", mergedFilePath)
	}

	downloader.bar = progressBar(stream.Size)
	if !downloader.option.Silent {
		// 创建并启动一个新的进度条
		downloader.bar.Start()
	}

	if len(stream.Parts) == 1 {
		// only one fragment
		var err error
		if downloader.option.MultiThread {
			err = downloader.multiThreadSave(stream.Parts[0], data.URL, mergedFilePath)
		} else {
			err = downloader.save(stream.Parts[0], data.URL, mergedFilePath)
		}

		if err != nil {
			return err
		}
		downloader.bar.Finish()
		return nil
	}

	return nil
}

func (downloader Downloader) multiThreadSave(stream *videer.Part, url, title string) error {

	return nil
}

func (downloader Downloader) save(part *videer.Part, url, filePath string) error {
	tempFilePath := filePath + DOWNLOAD_FILE_EXT
	tempFileSize, _, err := utils.FileSize(tempFilePath)
	if err != nil {
		return err
	}
	var (
		file      *os.File
		fileError error
	)
	if tempFileSize > 0 {
		headers := map[string]string{
			"Referer": url,
		}
		// range start from 0, 0-1023 means the first 1024 bytes of the file
		headers["Range"] = fmt.Sprintf("bytes=%d-", tempFileSize)
		file, fileError = os.OpenFile(tempFilePath, os.O_APPEND|os.O_WRONLY, 0644)
		downloader.bar.Add64(tempFileSize)
	} else {
		file, fileError = os.Create(tempFilePath)
		if fileError != nil {
			return fileError
		}
	}
	defer file.Close()

	headers := map[string]string{
		"Referer": url,
	}
	var start, end, chunkSize int64
	chunkSize = int64(downloader.option.ChunkSizeMB) * 1024 * 1024
	chunk := part.Size / chunkSize
	if part.Size%chunkSize != 0 {
		chunk += 1
	}

	var i int64 = 1
	for ; i <= chunk; i++ {
		// 1-999  1000 - 1999
		end = start + chunkSize - 1
		headers["Range"] = fmt.Sprintf("bytes=%d-%d", start, end)
		temp := start
		for {
			written, err := downloader.writeFile(file, part.URL, headers)
			if err == nil {
				break
			}
			temp += written
			headers["Range"] = fmt.Sprintf("bytes=%d-%d", temp, end)
			time.Sleep(1 * time.Second)
		}
		start = end + 1
	}

	return nil
}

func (downloader Downloader) writeFile(file *os.File, url string, headers map[string]string) (int64, error) {
	res, err := request.Request(http.MethodGet, url, nil, headers)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close() // nolint

	// 进度条更新
	barWriter := downloader.bar.NewProxyWriter(file)

	// 写入到文件中
	written, copyErr := io.Copy(barWriter, res.Body)
	if copyErr != nil && copyErr != io.EOF {
		return written, errors.Errorf("file copy error: %s", copyErr)
	}
	return written, nil
}
