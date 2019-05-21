package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"html"
	"net/http"

	_ "GoStuff/goi18n/catalog"
)

func PrintMessage01(w http.ResponseWriter, r *http.Request) {
	tag, _, _ := matcher.Match(language.MustParse("en"))
	printer := message.NewPrinter(tag)
	printer.Fprintf(w, "Hi, %v", html.EscapeString(r.Host))
}
