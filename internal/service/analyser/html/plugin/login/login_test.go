package login

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
	<form>
		<input type="password"></input>
	</form>
</body>
</html>`
const htm2 = `<!DOCTYPE html>
<html>
<head>
    <title>hello world</title>
</head>
<body>
    body content
	<p>more content</p>
	<form>
	<input type="password"/>
	<input type="password"/>
	</form>
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
	if !processor.HasLoginForm {
		t.Errorf("Should have login form Set:%v should be: %v", processor.HasLoginForm, true)
	}

	node, _ = html.Parse(strings.NewReader(htm2))
	traverse(node, processor)
	if !processor.HasLoginForm {
		t.Errorf("Should have login form even after parsing other form Set:%v should be: %v", processor.HasLoginForm, true)
	}

	processor = &Processor{}

	node, _ = html.Parse(strings.NewReader(htm2))
	traverse(node, processor)
	if processor.HasLoginForm {
		t.Errorf("Should not have login form Set:%v should be: %v", processor.HasLoginForm, false)
	}
}
