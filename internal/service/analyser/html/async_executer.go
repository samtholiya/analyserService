package html

import (
	"sync"

	"github.com/samtholiya/analyserService/internal/service/analyser/html/plugin"
	"golang.org/x/net/html"
)

// workerPoolCount worker routines count
const workerPoolCount int = 10

type asyncExecuterData struct {
	Processor plugin.ProcessorInterface
	Node      *html.Node
	WaitGroup *sync.WaitGroup
}

var (
	executerChannel = make(chan asyncExecuterData)
)

func init() {
	for i := 0; i < workerPoolCount; i++ {
		go asyncExecuter()
	}
}

func asyncExecuter() {
	for {
		executerData := <-executerChannel
		executerData.Processor.Execute(executerData.Node)
		executerData.WaitGroup.Done()
	}
}
