package anchor

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
	<a href="https://www.google.com"></a>
</body>
</html>`

func traverse(node *html.Node, processor *Processor) {
	processor.Execute(node)
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, processor)
	}
}

func TestAnchorExternalLink(t *testing.T) {
	processor := &Processor{
		url: "https://www.google.co.in/",
	}
	node, _ := html.Parse(strings.NewReader(htm))
	traverse(node, processor)
	if processor.Links != 1 {
		t.Errorf("Links is not set correctly Set:%v should be: %v", processor.Links, 1)
	}
	if processor.ExternalLinks != 1 {
		t.Errorf("ExternalLinks is not set correctly Set:%v should be: %v", processor.ExternalLinks, 1)
	}
	if processor.InternalLinks != 0 {
		t.Errorf("InternalLinks is not set correctly Set:%v should be: %v", processor.InternalLinks, 1)
	}

	if processor.AccessibleLinks != 1 {
		t.Errorf("AccessibleLinks is not set correctly Set:%v should be: %v", processor.AccessibleLinks, 1)
	}

	if processor.InaccessibleLinks != 0 {
		t.Errorf("InaccessibleLinks is not set correctly Set:%v should be: %v", processor.InaccessibleLinks, 0)
	}
}

const htm2 = `<!DOCTYPE html>
<html>
<head>
    <title>hello world</title>
</head>
<body>
    body content
	<p>more content</p>
	<a href="/images"></a>
</body>
</html>`

func TestAnchorInternalLink(t *testing.T) {
	processor := &Processor{
		url: "https://www.google.co.in/",
	}
	node, _ := html.Parse(strings.NewReader(htm2))

	traverse(node, processor)
	if processor.Links != 1 {
		t.Errorf("Links is not set correctly Set:%v should be: %v", processor.Links, 1)
	}
	if processor.ExternalLinks != 0 {
		t.Errorf("ExternalLinks is not set correctly Set:%v should be: %v", processor.ExternalLinks, 1)
	}
	if processor.InternalLinks != 1 {
		t.Errorf("InternalLinks is not set correctly Set:%v should be: %v", processor.InternalLinks, 1)
	}

	if processor.AccessibleLinks != 1 {
		t.Errorf("AccessibleLinks is not set correctly Set:%v should be: %v", processor.AccessibleLinks, 1)
	}

	if processor.InaccessibleLinks != 0 {
		t.Errorf("InaccessibleLinks is not set correctly Set:%v should be: %v", processor.InaccessibleLinks, 0)
	}
}
