package main

import (
	"fmt"
	"net/http"
)

func main() {

	prefix := "http://"
	url := "baidu.com"
	url = prefix + url
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", resp)
}
