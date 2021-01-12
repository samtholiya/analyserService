package anchor

import (
	"net/http"
	"strings"
	"sync/atomic"

	"github.com/samtholiya/analyserService/internal/service/analyser/html/plugin/helper"
	htmlNative "golang.org/x/net/html"
)

type Processor struct {
	url               string `json:"-"`
	Links             uint64
	InaccessibleLinks uint64
	AccessibleLinks   uint64
	InternalLinks     uint64
	ExternalLinks     uint64
}

func (p *Processor) Execute(node *htmlNative.Node) {

	isAnchor := node.Data == "a" && node.Type == htmlNative.ElementNode
	if !isAnchor {
		return
	}
	// Extract the href value, if there is one
	url, ok := helper.GetAttrValue(node, "href")
	if !ok {
		return
	}
	p.incrementLinkCounters(url)
	p.classifyLinks(url)
}

func (p *Processor) incrementLinkCounters(url string) {
	atomic.AddUint64(&p.Links, 1)

	hasProto := strings.Index(url, "http") == 0
	if hasProto {
		atomic.AddUint64(&p.ExternalLinks, 1)
		return
	}
	atomic.AddUint64(&p.InternalLinks, 1)
}

func (p *Processor) classifyLinks(url string) {
	genURL := url
	if !strings.HasPrefix(url, "http") {
		genURL = p.url + genURL
	}
	res, err := http.Get(genURL)
	isErrorPresent := err != nil
	isErrorStatusCode := res != nil && res.StatusCode >= 400
	isLinkAccessible := isErrorPresent || isErrorStatusCode
	if isLinkAccessible {
		atomic.AddUint64(&p.InaccessibleLinks, 1)
		return
	}
	atomic.AddUint64(&p.AccessibleLinks, 1)
}
