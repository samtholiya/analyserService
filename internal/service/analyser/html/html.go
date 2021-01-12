package html

import (
	"io"
	"net/http"
)

//Crawler analyses HTML documents
type Crawler struct {
}

// FromReader loads data from reader
func (h *Crawler) FromReader(url string, reader io.Reader) (map[string]interface{}, error) {
	return Parse(url, reader)
}

// FromURL loads data from url
func (h *Crawler) FromURL(url string) (map[string]interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return h.FromReader(url, res.Body)
}
