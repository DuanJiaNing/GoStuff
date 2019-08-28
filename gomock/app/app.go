package app

import (
	"../spider"
)

func GetSpiderVersion(s spider.Spider) string {
	return s.GetVersion()
}
