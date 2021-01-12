package helper

import (
	"testing"

	"golang.org/x/net/html"
)

func TestGetAttrValue(t *testing.T) {
	val, ok := GetAttrValue(&html.Node{
		Attr: []html.Attribute{
			html.Attribute{Key: "href", Val: "hello"},
		},
	}, "href")
	if !ok || val != "hello" {
		t.Errorf("Error in helper funtioin GetAttrValue positive test case val:%v ok:%v", val, ok)
	}
	val, ok = GetAttrValue(&html.Node{
		Attr: []html.Attribute{
			html.Attribute{Key: "href", Val: "hello"},
		},
	}, "hre")
	if ok || val != "" {
		t.Errorf("Error in helper funtioin GetAttrValue negative test case")
	}
}
