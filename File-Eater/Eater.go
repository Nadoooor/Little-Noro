package Eater

import (
	"os"
	// "log"
	"encoding/json"
	"main/Sounds"
)

type Progress struct {
	Points int `json:"points"`
	Count  int `json:"count"`
	Food   int `json:"food"`
	Water  int `json:"water"`
}

func save(p Progress) {
	data, err := os.OpenFile("progress.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
	}
	defer data.Close()
	err = json.NewEncoder(data).Encode(&p)
	if err != nil {
	}
}
func load() Progress {
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {
	}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {
	}
	return progress
}

func Loadcache() Progress {
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {
	}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {
	}
	return progress
}

func FileEater(File string) {
	var progress Progress = load()

	_ = os.Remove(File)
	if progress.Points >= 60 {
		Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
		progress.Points = progress.Points + 8
	} else {
		Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
		progress.Points += 5
	}
	save(progress)

}
func Currpoints() int {
	progress := load()
	return progress.Points
}
func CurrFood() int {
	progress := load()
	return progress.Food
}
func CurrWater() int {
	progress := load()
	return progress.Water
}

func levelsound() {
	var progress Progress = load()
	var points int = progress.Points
	if points < 20 {

	} else if points == 20 {
		Sounds.Sound("Level.mp3")
		progress.Food += 100

	} else if points == 40 {
		Sounds.Sound("Level.mp3")

		progress.Water += 100

	} else if points == 60 {
		Sounds.Sound("Level.mp3")
		progress.Food += 100

	} else if points == 130 {
		Sounds.Sound("Level.mp3")

		progress.Water += 100
	} else if points > 9999 {
		Sounds.Sound("Level.mp3")

	}
	save(progress)

}

func FeedSpammer() {
	var progress Progress = load()
	if progress.Food > 0 {
		progress.Food = progress.Food - 1
		progress.Count = progress.Count + 1
		if progress.Count >= 10 {
			levelsound()
			progress.Count = 0
			progress.Points++
			Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
		}
		save(progress)
	}
}

func DrinkSpammer() {
	var progress Progress = load()
	if progress.Water > 0 {
		progress.Water = progress.Water - 1
		progress.Count = progress.Count + 1

		if progress.Count >= 10 {
			levelsound()
			progress.Count = 0
			progress.Points++
			Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
		}

		save(progress)

	}

}
