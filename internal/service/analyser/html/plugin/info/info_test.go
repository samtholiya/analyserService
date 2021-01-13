package info

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const htm = `<!DOCTYPE html>
<html>
<head>
    <title>hello world</title>
</head>
<body>
    body content
    <p>more content</p>
</body>
</html>`

const htm2 = `<!DOCTYPE html>
<html>
<head>
</head>
<body>
    body content
    <p>more content</p>
</body>
</html>`

func traverse(node *html.Node, processor *Processor) {
	processor.Execute(node)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, processor)
	}
}

func TestTitleExecute(t *testing.T) {
	processor := &Processor{}
	node, _ := html.Parse(strings.NewReader(htm))
	traverse(node, processor)
	if processor.Title != "hello world" {
		t.Errorf("Title is not set correctly Set:%v should be: %v", processor.Title, "hello world")
	}
}

func TestTitleExecuteNegative(t *testing.T) {
	processor := &Processor{}
	node, _ := html.Parse(strings.NewReader(htm2))
	traverse(node, processor)
	if processor.Title != "" {
		t.Errorf("Title is not set correctly Set:%v should be: %v", processor.Title, "hello world")
	}
}

func TestGetPluginName(t *testing.T) {
	processor := New("mockurl")
	if processor.GetProcessorName() != "Info" {
		t.Errorf("Plugin name should be Info found %v", processor.GetProcessorName())
	}
}
