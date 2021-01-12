package title

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
	processor := &Processor{}
	if processor.GetProcessorName() != "Title" {
		t.Errorf("Plugin name should be Title found %v", processor.GetProcessorName())
	}
}

func TestGetPluginJSON(t *testing.T) {
	processor := &Processor{
		Title: "Title",
	}
	if data, err := processor.MarshalJSON(); err != nil || string(data) != "\"Title\"" {
		t.Errorf("MarshalJSON should return \"Title\" found: %v", string(data))
	}
}
