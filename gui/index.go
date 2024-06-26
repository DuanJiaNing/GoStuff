package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func index(w fyne.Window) fyne.CanvasObject {
	content := container.NewStack()
	title := widget.NewLabel("Title")
	intro := widget.NewLabel("Intro")
	intro.Wrapping = fyne.TextWrapWord
	split := container.NewHSplit(
		makeNav(func(t Tool) {
			title.SetText(t.Title)
			intro.SetText(t.Intro)
			content.Objects = []fyne.CanvasObject{t.View(w)}
			content.Refresh()
		}),
		container.NewBorder(container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content),
	)

	return split
}

func makeNav(setTool func(t Tool)) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return ToolIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := ToolIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Sub tools")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t := Tools[uid]
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := Tools[uid]; ok {
				setTool(t)
			}
		},
	}

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}

func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}
