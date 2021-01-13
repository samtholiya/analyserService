package info

import (
	"github.com/samtholiya/analyserService/internal/service/analyser/html/plugin"

	htmlNative "golang.org/x/net/html"
)

const name = "Info"

func init() {
	plugin.RegisterProcessor(name, New)
}

func (p *Processor) GetProcessorName() string {
	return name
}

func New(url string) plugin.ProcessorInterface {
	return &Processor{}
}

type Processor struct {
	Title string
}

func (p *Processor) Execute(node *htmlNative.Node) {
	isTitle := node.Data == "title" && node.Type == htmlNative.ElementNode
	if !isTitle {
		return
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == htmlNative.TextNode {
			p.Title = p.Title + c.Data
		}
	}
}
