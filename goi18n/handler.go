package main

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,el,zh

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"html"
	"net/http"

	_ "GoStuff/goi18n/catalog"
)

var matcher = language.NewMatcher(message.DefaultCatalog.Languages())

func PrintMessage(w http.ResponseWriter, r *http.Request) {
	tag, _, _ := matcher.Match(language.MustParse("en"))
	printer := message.NewPrinter(tag)
	printer.Fprintf(w, "Hello, %v", html.EscapeString(r.Host))
}
