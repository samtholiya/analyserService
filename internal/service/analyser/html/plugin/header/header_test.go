package header

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
	<p>more content <h1>hello world </h1> </p>
	<h1>hello world </h1>
</body>
</html>`

const htm2 = `<!DOCTYPE html>
<html>
<head>
</head>
<body>
    body content
	<p>more content</p>
	<h1>hello world </h1>
	<h2>hello world <h2>hello world </h2></h2>
	<h3>hello world </h3>
	
</body>
</html>`

func traverse(node *html.Node, processor *Processor) {
	processor.Execute(node)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, processor)
	}
}

func TestHeadersExecute(t *testing.T) {
	processor := New("http://something.com")
	node, _ := html.Parse(strings.NewReader(htm))
	traverse(node, processor.(*Processor))
	if processor.(*Processor).headers["h1"] != 2 {
		t.Errorf("Headers is not set correctly Set:%v should be: h1:2", processor.(*Processor).headers)
	}

	processor = New("http://something.com")
	node, _ = html.Parse(strings.NewReader(htm2))
	traverse(node, processor.(*Processor))
	if processor.(*Processor).headers["h1"] != 1 && processor.(*Processor).headers["h2"] != 2 && processor.(*Processor).headers["h3"] != 1 {
		t.Errorf("Title is not set correctly Set:%v should be: h1:1 h2:2 h3:1", processor.(*Processor).headers)
	}

}

func TestGetPluginName(t *testing.T) {
	processor := &Processor{}
	if processor.GetProcessorName() != "Headers" {
		t.Errorf("Plugin name should be Headers found %v", processor.GetProcessorName())
	}
}

func TestGetPluginJSON(t *testing.T) {
	processor := &Processor{
		headers: map[string]uint{
			"h1": 1,
		},
	}
	if data, err := processor.MarshalJSON(); err != nil || string(data) != "{\"h1\":1}" {
		t.Errorf("MarshalJSON should return {\"h1\":1} found: %v", string(data))
	}
}
