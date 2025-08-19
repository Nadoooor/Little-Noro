package main

import (
	"image/color"
	"main/Images"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	var bg *canvas.Image
	myApp := app.New()
	myWindow := myApp.NewWindow("Little-Noro")
	timen := time.Now().Format("3:04 PM")

	if strings.Contains(timen, "PM") {
		bg = canvas.NewImageFromFile("Back.png")
		bg.FillMode = canvas.ImageFillStretch
	} else {
		bg = canvas.NewImageFromFile("Backnight.png")
		bg.FillMode = canvas.ImageFillStretch
	}

	bottle, _ := fyne.LoadResourceFromPath("bottle.png")
	food, _ := fyne.LoadResourceFromPath("Foodk.png")
	plate := canvas.NewImageFromFile("plate.png")
	pet := canvas.NewImageFromImage(Images.Images())
	points := widget.NewLabel("points")
	points.SetText("Points:")
	foodtxt := widget.NewLabel("Food")
	foodtxt.SetText("Food:")
	water := widget.NewLabel("Water")
	water.SetText("Water:")
	timet := widget.NewLabel("Time")

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			fyne.Do(func() {
				timet.SetText("time: " + timen)
				timet.Refresh()
			})
		}
	}()

	black := canvas.NewRectangle(color.Black)

	button2 := widget.NewButtonWithIcon("fuck", food, func() {

	})

	button := widget.NewButtonWithIcon("spam", bottle, func() {
	})

	button2.Resize(fyne.NewSize(100, 35))
	button2.Move(fyne.NewPos(265, 475))
	pet.Resize(fyne.NewSize(300, 300))
	pet.Move(fyne.NewPos(50, 100))
	button.Resize(fyne.NewSize(100, 35))
	button.Move(fyne.NewPos(25, 475))
	plate.Resize(fyne.NewSize(200, 200))
	plate.Move(fyne.NewPos(95, 400))
	points.Resize(fyne.NewSize(100, 35))
	points.Move(fyne.NewPos(300, 0))
	foodtxt.Resize(fyne.NewSize(100, 35))
	foodtxt.Move(fyne.NewPos(100, 0))
	water.Resize(fyne.NewSize(100, 35))
	water.Move(fyne.NewPos(200, 0))
	timet.Resize(fyne.NewSize(100, 35))
	timet.Move(fyne.NewPos(0, 0))
	bg.Resize(fyne.NewSize(400, 600))
	black.Resize(fyne.NewSize(400, 35))
	black.Move(fyne.NewPos(0, 0))

	container1 := container.NewWithoutLayout(bg, button, button2, plate, pet, black, water, points, foodtxt, timet)
	myWindow.SetContent(container1)
	myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()
}
