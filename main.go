package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func create_url(question string) string {
	return fmt.Sprintf("https://www.google.com/search?client=opera&q=%s&sourceid=opera&ie=UTF-8&oe=UTF-8&bshm=rimc/1", url.QueryEscape(question))
}
func make_request() string {
	var question string = os.Args[1]
	res, err := http.Get(create_url(question))
	if err != nil {
		fmt.Println("Error making google request")
		os.Exit(1)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	return string(resBody)

}

func main() {
	html_string := make_request()
	parse_html(html_string)
}
