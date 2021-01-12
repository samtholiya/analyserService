package helper

import (
	"golang.org/x/net/html"
)

// GetAttrValue Helper function to pull the key attribute from a Token
func GetAttrValue(node *html.Node, key string) (value string, ok bool) {
	// Iterate over token attributes until we find an "href"
	for _, a := range node.Attr {
		if a.Key == key {
			value = a.Val
			ok = true
			return
		}
	}
	return
}
