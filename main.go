package main

import (
	"github.com/maxilian/panic-reloader/pkg/services/contexts"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()
var list = tview.NewList()
var pages = tview.NewPages()

func main() {
	pages.AddPage("mainmenu", list, true, true)

	pages.SwitchToPage("mainmenu")
	list.SetBorder(true)
	list.SetTitle("MAIN MENU")
	list.Clear()

	list.AddItem("Pilih Kubectl Context", "Menu untuk memilih cluster kubernetes", 'a', func() {
		v := contexts.NewKubeContext(app, pages)
		//app.SetRoot(v.Content(), true)
		v.GetContext()
	})

	list.AddItem("Refresh Kubernetes Secret", "Some explanatory text", 'b', func() {
		list.Clear()
		dropdown := tview.NewDropDown().
			SetLabel("Select an option (hit Enter): ").
			SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)

		app.SetRoot(dropdown, true).SetFocus(dropdown)
	})

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
