package main

import (
	"fmt"
	"main/Sounds"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"



	"main/File-Eater"
	// "github.com/gdamore/tcell/v2"
	// "os"
)


func main() {
var app *tview.Application = tview.NewApplication()
	var wm *winman.Manager= winman.NewWindowManager()
	
	var createForm = func(modal bool) *winman.WindowBase {

		var form *tview.Form= tview.NewForm()
		var All *tview.Grid= tview.NewGrid().
		SetSize(10,4,0,0)
		
		var window *winman.WindowBase= winman.NewWindow().
			SetRoot(All).
			SetResizable(true).
			SetDraggable(true).
			SetModal(modal)

		
		
		var display *tview.TextView = tview.NewTextView().
		SetText("").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)
		var Points *tview.TextView = tview.NewTextView().
		SetText("Points").
		SetTextAlign(tview.AlignLeft).SetSize(2,2)

			All.AddItem(form.AddInputField("Enter The FilePath", "", 40, nil, nil).
			AddFormItem(display).
			AddFormItem(Points).
			AddButton("Feed The File", func() {
				display.SetText(fmt.Sprintln(Eater.FileEater(form.GetFormItem(0).(*tview.InputField).GetText()), form.GetFormItem(0).(*tview.InputField).GetText()))
				Points.SetText(fmt.Sprintf("Points: %d", Eater.Currpoints()))
			}).
			AddButton("T 2 B", func() {
				Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
			}).
			AddButton("T 2 B64", func() {
				Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
			}).
			AddButton("T 2 B32", func() {
				Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
			}), 0,0,5,4,0,0,false)
		

		var title string= fmt.Sprintln("Little Noro:" + Eater.Levels())
		window.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)
		// window.SetRect(2, 2, 50, 30)


		


		
		wm.AddWindow(window)
		return window
	}

	for i := 0; i < 1; i++ {
		createForm(false).Show()
	}

	if err := app.SetRoot(wm, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	
}

	



