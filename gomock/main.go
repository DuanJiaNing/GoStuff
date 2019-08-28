package main

import (
	"./app"
	"./spider"
	"fmt"
)

func main() {
	v := app.GetSpiderVersion(spider.NewSpider())
	if v != "go1.8.3" {
		fmt.Println("Get wrong version", v)
	}
}
