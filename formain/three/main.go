package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"golang.org/x/text/language"
	"strings"
	"text/template"
)

func main1() {
	//str := "{{.Z}}/{{.X}}/{{.Y}}/{{.YEAR}}/{{.DATE}}.jpg"
	str := "{{.Z}}/{{.X}}/{{.Y}}/map_tile.png"
	tmpl, _ := template.New("path_template").Parse(str)
	buffer := bytes.NewBuffer([]byte{})
	const truncatePart string = "XXX"
	tmpl.Execute(buffer, map[string]interface{}{
		"X":    12,
		"Y":    23,
		"Z":    11,
		"YEAR": 2018,
	})

	s := buffer.String()
	fmt.Println(strings.TrimSuffix(s, truncatePart+".jpg"))
	fmt.Println(s)
	fmt.Println(s[:strings.LastIndex(s, "/")+1])
}

func main() {
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("AR/OES/HENDERSON.DAIREAUX/RAN/LAPAZ/duan_test_0001")))
}

func main2() {
	matcher := language.NewMatcher([]language.Tag{
		//language.Make("es-AR"),
		language.AmericanEnglish,
	})
	matched, _, _ := matcher.Match(language.Make(""))
	fmt.Println(matched)
}
