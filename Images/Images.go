package Images

import (
	"log"
	"os"

	"encoding/json"

	"fyne.io/fyne/v2"
)

type Progress struct {
	Points int `json:"points"`
	Count  int `json:"count"`
	Food   int `json:"food"`
	Water  int `json:"water"`
}

func Images() fyne.Resource {
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {
	}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {
	}
	points := progress.Points
	if points < 20 {
		img, err := fyne.LoadResourceFromPath("Little-Noro.png")
		if err != nil {
			log.Println("Failed to load image:", err)
			return nil
		}

		return img
	} else if points >= 20 && points < 40 {
		img, err := fyne.LoadResourceFromPath("Noro.png")
		if err != nil {
			log.Println("Failed to load image:", err)
			return nil
		}

		return img
	} else if points >= 40 && points < 60 {
		img, err := fyne.LoadResourceFromPath("Hero_Noro.png")
		if err != nil {
			log.Println("Failed to load image:", err)
			return nil
		}

		return img
	} else if points >= 60 && points < 9999 {
		img, err := fyne.LoadResourceFromPath("Noronium.png")
		if err != nil {
			log.Println("Failed to load image:", err)
			return nil
		}

		return img
	} else if points >= 9999 {
		img, err := fyne.LoadResourceFromPath("Cheater.png")
		if err != nil {
			log.Println("Failed to load image:", err)
			return nil
		}

		return img
	} else {
		img, err := fyne.LoadResourceFromPath("Notfound.png")
		if err != nil {
			log.Println("Failed to load image:", err)
			return nil
		}

		return img
	}

}
