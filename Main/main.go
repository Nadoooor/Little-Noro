package main

import (
	"image/color"
	"log"
	Eater "main/File-Eater"
	"main/Images"
	"runtime"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func checkMemoryLimit(maxMB uint64) bool {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	usedMB := m.Alloc / 1024 / 1024
	return usedMB < maxMB
}

func main() {

	var bg *canvas.Image
	var myApp fyne.App = app.New()
	var myWindow fyne.Window = myApp.NewWindow("Little-Noro")
	var timen string = time.Now().Format("3:04 PM")
	rebgnight, _ := fyne.LoadResourceFromPath("Backnight.png")
	rebg, _ := fyne.LoadResourceFromPath("Back.png")
	pla, _ := fyne.LoadResourceFromPath("plate.png")

	if strings.Contains(timen, "PM") {
		bg = canvas.NewImageFromResource(rebg)
	} else {
		bg = canvas.NewImageFromResource(rebgnight)
	}

	bottle, _ := fyne.LoadResourceFromPath("bottle.png")

	food, _ := fyne.LoadResourceFromPath("Foodk.png")

	var plate *canvas.Image = canvas.NewImageFromResource(pla)
	var pet *canvas.Image = canvas.NewImageFromResource(Images.Images())
	var points *widget.Label = widget.NewLabel("points")
	points.SetText("Points:")
	var foodtxt *widget.Label = widget.NewLabel("Food")
	foodtxt.SetText("Food:")
	var water *widget.Label = widget.NewLabel("Water")
	water.SetText("Water:")
	var timet *widget.Label = widget.NewLabel("Time")

	go func() {
		var ticker *time.Ticker = time.NewTicker(1 * time.Second)
		var lastpet fyne.Resource
		lastpet = pet.Resource
		if checkMemoryLimit(75) {
			for range ticker.C {
				fyne.Do(func() {
					var timen string = time.Now().Format("3:04 PM")
					newpet := Images.Images()
					if lastpet != newpet {
						pet.Resource = Images.Images()
						pet.Refresh()
						lastpet = pet.Resource
					}
					if timet.Text != timen {
						timet.SetText("time: " + timen)
						timet.Refresh()
					}
					progress := Eater.Loadcache()
					if points.Text != ("Points: " + strconv.Itoa(progress.Points)) {
						points.SetText("Points: " + strconv.Itoa(Eater.Currpoints()))
					}

					if foodtxt.Text != ("Food: " + strconv.Itoa(progress.Food)) {
						foodtxt.SetText("Food: " + strconv.Itoa(Eater.CurrFood()))
					}
					if water.Text != ("Water: " + strconv.Itoa(progress.Water)) {
						water.SetText("Water: " + strconv.Itoa(Eater.CurrWater()))
					}

				})
			}
		} else {
			log.Println("Memory limit reached—skipping update")
		}

	}()

	var black *canvas.Rectangle = canvas.NewRectangle(color.Black)

	var button2 *widget.Button = widget.NewButtonWithIcon("spam", food, func() {
		Eater.FeedSpammer()
	})

	var button *widget.Button = widget.NewButtonWithIcon("spam", bottle, func() {
		Eater.DrinkSpammer()
	})

	button2.Resize(fyne.NewSize(100, 35))
	button2.Move(fyne.NewPos(265, 465))
	pet.Resize(fyne.NewSize(300, 300))
	pet.Move(fyne.NewPos(50, 100))
	button.Resize(fyne.NewSize(100, 35))
	button.Move(fyne.NewPos(25, 465))
	plate.Resize(fyne.NewSize(200, 200))
	plate.Move(fyne.NewPos(95, 390))
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

	var container1 *fyne.Container = container.NewWithoutLayout(bg, button, button2, plate, pet, black, water, points, foodtxt, timet)
	myWindow.SetContent(container1)

	myWindow.SetOnDropped(func(pos fyne.Position, uris []fyne.URI) {
		for _, uri := range uris {
			Eater.FileEater(uri.Path())
		}
	})

	myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()
}
