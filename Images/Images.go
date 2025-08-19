package Images

import (
	"image"
	"image/png"

	"os"

	"encoding/json"
)

type Progress struct {
	Points int `json:"points"`
	Count  int `json:"count"`
	Food   int `json:"food"`
	Water  int `json:"water"`
}

func Images() image.Image {
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
		file, err := os.Open("Little-Noro.png")
		if err != nil {
		}

		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
		}

		return img
	} else if points >= 20 && points < 40 {
		file, err := os.Open("Noro.png")
		if err != nil {
		}

		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
		}

		return img
	} else if points >= 40 && points < 60 {
		file, err := os.Open("Hero_Noro.png")
		if err != nil {
		}

		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
		}

		return img
	} else if points >= 60 && points < 9998 {
		file, err := os.Open("Noronium.png")
		if err != nil {
		}

		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
		}

		return img
	} else if points >= 9999 {
		file, err := os.Open("Cheater.png")
		if err != nil {
		}

		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
		}

		return img
	} else {
		file, err := os.Open("Notfound.png")
		if err != nil {
		}

		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
		}

		return img

	}

}
