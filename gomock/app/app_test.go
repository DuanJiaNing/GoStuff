package app

import (
	"../spider"
	"testing"
)

func Test_GetGoVersion(t *testing.T) {
	v := GetSpiderVersion(spider.NewSpider())
	if v != "go1.8.3" {
		t.Errorf("Get wrong version %s", v)
	}
}
