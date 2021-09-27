package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var content *widget.Entry

func setNote(n *note) {
	content.SetText(n.content)
}

func loadUI(l *notelist) fyne.CanvasObject {
	list := widget.NewVBox()
	for _, n := range l.notes {
		note := n
		item := widget.NewButton(n.title(), func() {
			setNote(note)
		})
		list.Append(item)
	}

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			note := l.add()
			setNote(note)
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)
	content = widget.NewMultiLineEntry()
	if len(l.notes) > 0 {
		content.SetText(l.notes[0].content)
	}

	left := fyne.NewContainerWithLayout(layout.NewBorderLayout(bar, nil, nil, nil),
		bar, list)

	ui := container.NewHSplit(left, content)
	ui.Offset = 0.25
	return ui
}

func main() {
	a := app.New()
	w := a.NewWindow("Rabbitaxi")
	a.Preferences()
	notes := &notelist{[]*note{
		{content: "Taxi 1  \n has 1 request"},
		{content: "Taxi 2  \n  has 2 requests."},
	}}
	w.SetContent(loadUI(notes))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
