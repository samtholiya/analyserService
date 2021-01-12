package html

import (
	"encoding/json"
	"testing"
)

func TestHTMLLoad(t *testing.T) {
	crawl := Crawler{}
	mapData, _ := crawl.FromURL("https://www.google.co.in")
	data, _ := json.Marshal(mapData)
	// TODO: html testing should have a particular testcases (Kind of integration testing)
	t.Log(string(data))
}
