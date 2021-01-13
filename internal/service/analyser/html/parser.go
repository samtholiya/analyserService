package html

import (
	"io"
	"sync"

	"github.com/samtholiya/analyserService/internal/service/analyser/html/plugin"
	"golang.org/x/net/html"

	// Plugins for parser
	_ "github.com/samtholiya/analyserService/internal/service/analyser/html/plugin/anchor"
	_ "github.com/samtholiya/analyserService/internal/service/analyser/html/plugin/form"
	_ "github.com/samtholiya/analyserService/internal/service/analyser/html/plugin/header"
	_ "github.com/samtholiya/analyserService/internal/service/analyser/html/plugin/info"
)

func traverse(waitGroup *sync.WaitGroup, node *html.Node, freshProcessors []plugin.ProcessorInterface) {
	for index := range freshProcessors {
		waitGroup.Add(1)
		executerChannel <- asyncExecuterData{
			Processor: freshProcessors[index],
			Node:      node,
			WaitGroup: waitGroup,
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverse(waitGroup, c, freshProcessors)
	}
}

//Parse Extract all http analytics from a given webpage
func Parse(url string, reader io.Reader) (map[string]interface{}, error) {

	node, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	freshProcessors := make([]plugin.ProcessorInterface, len(plugin.Processors))
	i := 0
	for key := range plugin.Processors {
		freshProcessors[i] = plugin.Processors[key](url)
		i++
	}

	waitGroup := &sync.WaitGroup{}
	traverse(waitGroup, node, freshProcessors)
	waitGroup.Wait()

	parsedData := make(map[string]interface{})
	for index := range freshProcessors {
		parsedData[freshProcessors[index].GetProcessorName()] = freshProcessors[index]
	}
	return parsedData, nil
}
