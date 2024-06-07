package request

import (
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// Request base request
func Request(method, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DisableCompression:  true,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Minute,
		Jar:       jar,
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if _, ok := headers["Referer"]; !ok {
		req.Header.Set("Referer", url)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Get request
func Get(url, refer string, headers map[string]string) (string, error) {
	body, err := GetByte(url, refer, headers)
	return string(body), err
}

// GetByte get request
func GetByte(url, refer string, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{}
	}
	if refer != "" {
		headers["Referer"] = refer
	}
	res, err := Request(http.MethodGet, url, nil, headers)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close() // nolint

	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(res.Body)
	case "deflate":
		reader = flate.NewReader(res.Body)
	default:
		reader = res.Body
	}
	defer reader.Close() // nolint

	body, err := io.ReadAll(reader)
	if err != nil && err != io.EOF {
		return nil, errors.WithStack(err)
	}
	return body, nil
}
