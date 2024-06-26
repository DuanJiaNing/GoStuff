package tool

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// WidgetScreen shows a panel containing widget demos
func WidgetScreen(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox(
		widget.NewLabel("Labels"),
		widget.NewButtonWithIcon("Icons", theme.HomeIcon(), func() {}),
		widget.NewSlider(0, 1))
	return container.NewCenter(content)
}

func MakeAccordionTab(_ fyne.Window) fyne.CanvasObject {
	link, err := url.Parse("https://fyne.io/")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	ac := widget.NewAccordion(
		widget.NewAccordionItem("A", widget.NewHyperlink("One", link)),
		widget.NewAccordionItem("B", widget.NewLabel("Two")),
		&widget.AccordionItem{
			Title:  "C",
			Detail: widget.NewLabel("Three"),
		},
	)
	ac.MultiOpen = true
	ac.Append(widget.NewAccordionItem("D", &widget.Entry{Text: "Four"}))
	return ac
}

func MakeButtonTab(_ fyne.Window) fyne.CanvasObject {
	disabled := widget.NewButton("Disabled", func() {})
	disabled.Disable()

	shareItem := fyne.NewMenuItem("Share via", nil)
	shareItem.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Twitter", func() { fmt.Println("context menu Share->Twitter") }),
		fyne.NewMenuItem("Reddit", func() { fmt.Println("context menu Share->Reddit") }),
	)
	menuLabel := newContextMenuButton("tap me for pop-up menu with submenus", fyne.NewMenu("",
		fyne.NewMenuItem("Copy", func() { fmt.Println("context menu copy") }),
		shareItem,
	))

	return container.NewVScroll(container.NewVBox(
		widget.NewButton("Button (text only)", func() { fmt.Println("tapped text button") }),
		widget.NewButtonWithIcon("Button (text & leading icon)", theme.ConfirmIcon(), func() { fmt.Println("tapped text & leading icon button") }),
		&widget.Button{
			Alignment: widget.ButtonAlignLeading,
			Text:      "Button (leading-aligned, text only)",
			OnTapped:  func() { fmt.Println("tapped leading-aligned, text only button") },
		},
		&widget.Button{
			Alignment:     widget.ButtonAlignTrailing,
			IconPlacement: widget.ButtonIconTrailingText,
			Text:          "Button (trailing-aligned, text & trailing icon)",
			Icon:          theme.ConfirmIcon(),
			OnTapped:      func() { fmt.Println("tapped trailing-aligned, text & trailing icon button") },
		},
		disabled,
		&widget.Button{
			Text:       "Primary button",
			Importance: widget.HighImportance,
			OnTapped:   func() { fmt.Println("high importance button") },
		},
		&widget.Button{
			Text:       "Danger button",
			Importance: widget.DangerImportance,
			OnTapped:   func() { fmt.Println("tapped danger button") },
		},
		&widget.Button{
			Text:       "Warning button",
			Importance: widget.WarningImportance,
			OnTapped:   func() { fmt.Println("tapped warning button") },
		},
		layout.NewSpacer(),
		layout.NewSpacer(),
		menuLabel,
		layout.NewSpacer(),
	))
}

type contextMenuButton struct {
	widget.Button
	menu *fyne.Menu
}

func (b *contextMenuButton) Tapped(e *fyne.PointEvent) {
	widget.ShowPopUpMenuAtPosition(b.menu, fyne.CurrentApp().Driver().CanvasForObject(b), e.AbsolutePosition)
}

func newContextMenuButton(label string, menu *fyne.Menu) *contextMenuButton {
	b := &contextMenuButton{menu: menu}
	b.Text = label

	b.ExtendBaseWidget(b)
	return b
}
