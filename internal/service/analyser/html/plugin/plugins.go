package plugin

import (
	"log"
	"sync"

	"golang.org/x/net/html"
)

type ProcessorGenerator func(url string) ProcessorInterface

type ProcessorInterface interface {
	Execute(token *html.Node)
	GetProcessorName() string
}

var (
	processorsMu sync.RWMutex
	Processors   = make(map[string]ProcessorGenerator)
)

func RegisterProcessor(name string, processor ProcessorGenerator) {
	processorsMu.Lock()
	defer processorsMu.Unlock()
	if Processors == nil {
		panic("Processor: Register processor is nil")
	}
	if _, dup := Processors[name]; dup {
		panic("Processor: Register called twice for processor " + name)
	}
	log.Printf("Registered %v processor", name)
	Processors[name] = processor
}
