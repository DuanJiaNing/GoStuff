package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	fmt.Println(test(""))
	fmt.Println(test2(""))
}

func test(adminLevel string) string {
	parts := strings.Split(adminLevel, "/")
	parts = parts[:len(parts)-1]
	return strings.Join(parts, "/")
}

func test2(adminLevel string) string {
	if strings.IndexAny(adminLevel,"/") == -1 {
		return ""
	}
	return path.Dir(adminLevel)
}
