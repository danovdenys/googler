package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func find_links(node *html.Node) {
	var link string = node.Attr[0].Val
	if strings.HasPrefix(link, "/url?q=") && !strings.Contains(link, "google.com") {
		fmt.Println(strings.Replace(link, "/url?q=", "https://google.com/url?q=", -1))
	}
}
func resolve_html_tree(node *html.Node) {
	if node.Data == "a" {
		find_links(node)
	}
	if node.FirstChild == nil {
		return
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		resolve_html_tree(child)
	}
}

func parse_html(html_string string) {
	html_reader := strings.NewReader(html_string)
	node, _ := html.Parse(html_reader)
	resolve_html_tree(node)
}
