package main

import (
	"gui/tool"

	"fyne.io/fyne/v2"
)

type Tool struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var ToolIndex = map[string][]string{
	"":        {"welcome", "canvas", "widgets"},
	"widgets": {"accordion", "button"},
}

var Tools = map[string]Tool{
	"welcome":   {"Welcome", "", tool.WelcomeScreen},
	"canvas":    {"Canvas", "See the canvas capabilities.", tool.CanvasScreen},
	"widgets":   {"Widgets", "In this section you can see the features available in the toolkit widget set.\nExpand the tree on the left to browse the individual tutorial elements.", tool.WidgetScreen},
	"accordion": {"Accordion", "Expand or collapse content panels.", tool.MakeAccordionTab},
	"button":    {"Button", "Simple widget for user tap handling.", tool.MakeButtonTab},
}
