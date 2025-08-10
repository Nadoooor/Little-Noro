package main

import (
	"fmt"
	"main/Sounds"

	"github.com/epiclabs-io/winman"
	"github.com/rivo/tview"


	"image/png"
	// "github.com/gdamore/tcell/v2"
	"os"
)


func main() {
var app *tview.Application = tview.NewApplication()
	var wm *winman.Manager= winman.NewWindowManager()
	file, err := os.Open("Health.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	

	var createForm = func(modal bool) *winman.WindowBase {

		var form *tview.Form= tview.NewForm()
		var All *tview.Grid= tview.NewGrid().
		SetSize(10,4,0,0)
		
		var window *winman.WindowBase= winman.NewWindow().
			SetRoot(All).
			SetResizable(false).
			SetDraggable(true).
			SetModal(modal)

		
		
		var display *tview.TextView = tview.NewTextView().
		SetText("Output Here:").
		SetTextAlign(tview.AlignLeft)
		
			All.AddItem(form.AddInputField("Enter Here", "", 40, nil, nil).
			AddImage("test",img,,10,16777216).
			AddInputField("Enter ROT (Caesar)", "", 40, nil, nil).
			AddFormItem(display).
			AddButton("T 2 H", func() {
			Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
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

		
		


			
			

		var title string= fmt.Sprintln("Crypto-Hacks")
		window.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)
		window.SetRect(2, 2, 50, 30)
		window.AddButton(&winman.Button{
			Symbol:    'X',
			Alignment: winman.ButtonLeft,

		})

		var maxMinButton *winman.Button
		maxMinButton = &winman.Button{
			Symbol:    '▴',
			Alignment: winman.ButtonRight,
			OnClick: func() {
				if window.IsMaximized() {
					window.Restore()
					maxMinButton.Symbol = '▴'
				} else {
					window.Maximize()
					maxMinButton.Symbol = '▾'
				}
			},
		}
		window.AddButton(maxMinButton)
		
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

	



