package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("fyne.duanjn.com")
	logLifecycle(a)
	w := a.NewWindow("Tools")
	w.SetMaster()
	w.SetContent(index(w))
	w.Resize(fyne.NewSize(640, 460))

	w.ShowAndRun()
}
