package main

import (
	ui "github.com/VladimirMarkelov/clui"
	"github.com/nsf/termbox-go"
	"golang.org/x/crypto/ssh/terminal"
)

var execLog *ui.TextView
var list *ui.ListBox


func main() {
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	width, height, err := terminal.GetSize(0)
	if err != nil { panic(err) }

	view := ui.AddWindow(0, 0, width, height, "Cosmo")

	list = ui.CreateListBox(view,0,0,1)
	list.SetTitle("Actions")
	list.SetBackColor(ui.ColorBlack)
	list.SetTextColor(ui.ColorGreen)
	list.SetActiveTextColor(ui.ColorGreenBold)


	execLog = ui.CreateTextView(view, 0,0,1)
	execLog.SetTitle("Execution Log")
	execLog.SetBackColor(ui.ColorBlack)

	execLog.SetAutoScroll(true)

	actions := map[string]func(){
		"Install Dotfiles": func() {

		},
		"Open Dashboard": func() {
			writeLog("Not yet implemented, but it'll be totally rad man")

			window1 := ui.AddWindow(0,0,10,10, "buttons")

			ui.CreateButton(window1,0,0,"1",1)
			ui.CreateButton(window1,0,0,"2",1)
			ui.CreateButton(window1,0,0,"3",1)


		},

		"Show a bloody dialog mate": func() {
			ui.CreateAlertDialog("WOAH DUDE", "YO THIS IS SO RAD", "OKAY!")
		},

		"Quit": func() {
			ui.Stop()
		},
	}

	for k := range actions {
		list.AddItem(k)
	}
	list.SelectItem(0)

	list.OnKeyPress(func(key termbox.Key) bool {
		if key == termbox.KeyCtrlC || key == termbox.KeyEsc {
			ui.Stop()
		}

		if key == termbox.KeyEnter {
			for k,f := range actions {
				if list.SelectedItemText() == k {
					f()
				}
			}
		}
		return false
	})

	ui.MainLoop()

}

func ListEnabled(b bool) {
	if b {
		list.SetEnabled(b)
		list.SetActiveTextColor(ui.ColorGreenBold)
		list.SetTextColor(ui.ColorGreen)
		list.SetBackColor(ui.ColorBlack)
	} else {
		list.SetEnabled(b)
		list.SetActiveTextColor(ui.ColorBlack)
		list.SetTextColor(ui.ColorBlack)
		list.SetBackColor(ui.ColorWhite)
	}
}

func writeLog(in string) {
	execLog.AddText([]string{in})
	ui.RefreshScreen()
}
