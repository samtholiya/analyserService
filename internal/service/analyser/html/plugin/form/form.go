package form

import (
	"sync"
	"sync/atomic"

	"github.com/samtholiya/analyserService/internal/service/analyser/html/plugin"
	"github.com/samtholiya/analyserService/internal/service/analyser/html/plugin/helper"
	htmlNative "golang.org/x/net/html"
)

const name = "Form"

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
	HasLoginForm       bool
	passwordFieldCount uint64     `json:"-"`
	mutex              sync.Mutex `json:"-"`
}

func (p *Processor) Execute(node *htmlNative.Node) {
	isForm := node.Data == "form" && node.Type == htmlNative.ElementNode
	if !isForm {
		return
	}
	p.mutex.Lock()
	if p.HasLoginForm {
		return
	}
	p.traverse(node)
	p.HasLoginForm = p.passwordFieldCount == 1
	p.passwordFieldCount = 0
	p.mutex.Unlock()
}

func (p *Processor) traverse(node *htmlNative.Node) {
	isInput := node.Data == "input" && node.Type == htmlNative.ElementNode
	if !isInput {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			p.traverse(c)
		}
		return
	}
	val, ok := helper.GetAttrValue(node, "type")
	if !ok {
		return
	}
	isPasswordType := val == "password"
	if isPasswordType {
		atomic.AddUint64(&p.passwordFieldCount, 1)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		p.traverse(c)
	}
}
