package contexts

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type View interface {
	Content() tview.Primitive
}

type KubectlContext struct {
	pages *tview.Pages
	list  *tview.List
	app   *tview.Application
	// flex  *tview.Flex
}

func (kc *KubectlContext) Content() tview.Primitive {
	return kc.list
	// return kc.flex
}

func NewKubeContext(app *tview.Application, pages *tview.Pages) *KubectlContext {

	kc := &KubectlContext{
		pages: pages,
		list:  tview.NewList(),
		app:   app,
	}

	kc.pages.AddPage("kubecontext", kc.list, true, true)
	kc.pages.SwitchToPage("kubecontext")
	kc.list.SetTitle("PILIH KUBERNETES CONTEXT")
	kc.list.SetSelectedFunc(kc.onSelectContext)
	kc.list.SetBorder(true)

	kc.list.SetDoneFunc(func() {
		kc.pages.SwitchToPage("kubecontext")
	})

	kc.list.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		if e.Key() == tcell.KeyRune {
			switch e.Rune() {

			case 'q':
				kc.app.Stop()
			}
		}
		return e
	})

	//kc.app.SetFocus(kc.list)

	return kc

}

func (kc *KubectlContext) GetContext() {
	kc.list.Clear()
	config := getConfig()

	for ndx, context := range config.Contexts {
		var key rune
		if ndx < 10 {
			key = '0' + rune(ndx)
		}

		kc.list.AddItem(context.Name, "", key, nil)
		if context.Name == config.CurrentContext {
			kc.list.SetCurrentItem(ndx)
			kc.list.SetItemText(ndx, context.Name, "cluster saat ini")

		}
	}

	kc.list.AddItem("Quit", "Press q to exit", 'q', func() {
		kc.app.Stop()
	})

	kc.app.SetFocus(kc.list)
}

func (kc *KubectlContext) onSelectContext(i int, context string, _ string, cc rune) {
	//fmt.Print(cc)
	if context != "Quit" {
		useContext(context)
		kc.GetContext()
	}

}
