package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func main() {
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
	fmt.Println(s[:strings.LastIndex(s,"/")+1])
}
