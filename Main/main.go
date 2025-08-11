package main

import (
	"fmt"
	"main/Images"

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
		var form2 *tview.Form= tview.NewForm()
		var form3 *tview.Form= tview.NewForm()
		var form4 *tview.Form= tview.NewForm()
		var form5 *tview.Form= tview.NewForm()
		var form6 *tview.Form= tview.NewForm()
		var form7 *tview.Form= tview.NewForm()
		var All *tview.Grid= tview.NewGrid().
		SetSize(11,5,0,0)
		
		var window *winman.WindowBase= winman.NewWindow().
			SetRoot(All).
			SetResizable(false).
			SetDraggable(true).
			SetModal(modal)

		var log *tview.TextView = tview.NewTextView().
		SetTextAlign(tview.AlignLeft).SetSize(10, 10)

		var Points *tview.TextView = tview.NewTextView().
		SetText(fmt.Sprintf("Points: %d", Eater.Currpoints())).
		SetTextAlign(tview.AlignLeft).SetSize(10,10)
		var Food *tview.TextView = tview.NewTextView().
		SetText(fmt.Sprintf("Food: %d", Eater.CurrFood())).
		SetTextAlign(tview.AlignLeft).SetSize(10,10)
		var Water *tview.TextView = tview.NewTextView().
		SetText(fmt.Sprintf("Water: %d", Eater.CurrWater())).
		SetTextAlign(tview.AlignLeft).SetSize(10,10)

			All.AddItem(form.AddInputField("Enter The FullFilePath", "", 40, nil, nil), 0,0,2,5	,1,1,false).
			AddItem(form2.AddFormItem(Images.Images()), 3, 2, 2, 4, 1, 1, false).
			AddItem(form3.AddFormItem(Points), 10, 0, 1, 1, 0, 1, false).
			AddItem(form6.AddFormItem(Food), 2, 0, 1, 1, 0, 1, false).
			AddItem(form7.AddFormItem(Water), 6, 0, 1, 1, 0, 1, false).
			AddItem(form4.AddButton("Feed With File", func(){
				Eater.FileEater(form.GetFormItem(0).(*tview.InputField).GetText())
				Points.SetText(fmt.Sprintf("Points: %d", Eater.Currpoints()))
			}).AddButton("FeedSpam", func() {
				Eater.FeedSpammer()
				log.SetText(fmt.Sprintf("count %d", Eater.FeedSpammer()))
				Food.SetText(fmt.Sprintf("Food: %d", Eater.CurrFood()))
				Points.SetText(fmt.Sprintf("Points: %d", Eater.Currpoints()))
			}).AddButton("DrinkSpam", func() {
				Eater.DrinkSpammer()
				log.SetText(fmt.Sprintf("count %d", Eater.DrinkSpammer()))
				Water.SetText(fmt.Sprintf("Water: %d", Eater.CurrWater()))
				Points.SetText(fmt.Sprintf("Points: %d", Eater.Currpoints()))
			}), 10, 1, 1, 4, 1, 1, false).
			AddItem(form5.AddFormItem(log), 11, 0, 2, 1, 0, 1, false)

		var title string= fmt.Sprintln("Little Noro:" + Eater.Levels())
		window.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignCenter)
		window.SetRect(2, 2, 70, 45)
				 
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
