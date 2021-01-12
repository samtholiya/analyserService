package anchor

import "github.com/samtholiya/analyserService/internal/service/analyser/html/plugin"

const name = "Anchor"

func init() {
	plugin.RegisterProcessor(name, New)
}

func (p *Processor) GetProcessorName() string {
	return name
}

func New(url string) plugin.ProcessorInterface {
	return &Processor{
		url: url,
	}
}
