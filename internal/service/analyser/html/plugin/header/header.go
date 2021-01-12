package header

import (
	"encoding/json"
	"regexp"
	"sync"

	"github.com/samtholiya/analyserService/internal/service/analyser/html/plugin"
	htmlNative "golang.org/x/net/html"
)

const name = "Headers"

var regex *regexp.Regexp

func init() {
	plugin.RegisterProcessor(name, New)
	var err error
	regex, err = regexp.Compile("h\\d+")
	if err != nil {
		panic(err)
	}
}

func (p *Processor) GetProcessorName() string {
	return name
}

func New(url string) plugin.ProcessorInterface {
	return &Processor{
		headers: make(map[string]uint),
	}
}

type Processor struct {
	headers   map[string]uint
	headersMu sync.RWMutex
}

func (p *Processor) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.headers)
}

func (p *Processor) Execute(node *htmlNative.Node) {
	isHeader := regex.Match([]byte(node.Data)) && node.Type == htmlNative.ElementNode
	if !isHeader {
		return
	}
	p.headersMu.Lock()
	p.headers[node.Data]++
	p.headersMu.Unlock()
}
