package parser

import (
	"golang.org/x/net/html"
)

// FindAttr finds attr from node.
func FindAttr(node html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
