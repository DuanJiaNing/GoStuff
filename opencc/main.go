package main

import (
	"fmt"
	"log"

	"github.com/longbridgeapp/opencc"
)

/*
https://github.com/longbridgeapp/opencc

t2s.json 繁體到簡體
s2tw.json 簡體到臺灣正體
tw2s.json 臺灣正體到簡體
s2hk.json 簡體到香港繁體
hk2s.json 香港繁體到簡體
s2twp.json 簡體到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙
tw2sp.json 繁體（臺灣正體標準）到簡體並轉換爲中國大陸常用詞彙
t2tw.json 到臺灣正體
hk2t.json 香港繁體到繁體（OpenCC 標準）
t2hk.json 繁體（OpenCC 標準）到香港繁體
t2jp.json 繁體（OpenCC 標準，舊字體）到日文新字體
jp2t.json 日文新字體到繁體（OpenCC 標準，舊字體）
tw2t.json 臺灣正體到繁體（OpenCC 標準）
s2hk-finance.json 针对香港市场金融数据，做了特殊补充。
*/

func main() {
	s2t, err := opencc.New("s2t")
	if err != nil {
		log.Fatal(err)
	}

	in := `自然语言处理是人工智能领域中的一个重要方向。`
	out, err := s2t.Convert(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%s\n", in, out)
	//自然语言处理是人工智能领域中的一个重要方向。
	//自然語言處理是人工智能領域中的一個重要方向。
}
