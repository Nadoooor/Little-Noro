package Eater

import(
	"os"
	// "log"
	"encoding/json"
	"main/Sounds"

)
type Progress struct {
	Points int `json:"points"`
	Count  int `json:"count"`
	Food int `json:"food"`
	Water int `json:"water"`
}

func FileEater(File string) string{
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
       if err != nil {
		return err.Error()
	   }
	err = os.Remove(File)

	if err != nil {
		return "not found"
	}else if progress.Points >= 60 {
	Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
	
	data2, err := os.Open("progress.json")
	 if err != nil {
		return err.Error()
	   }
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
 if err != nil {
		return err.Error()
	   }
progress.Points = progress.Points + 8
	data, err := os.OpenFile("progress.json", os.O_WRONLY|os.O_CREATE, 0644)
	 if err != nil {
		return err.Error()
	   }
	defer data.Close()
	err = json.NewEncoder(data).Encode(&progress)
 if err != nil {
		return err.Error()
	   }
	   
return "+8"
	   
}else{
	Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
	
	data2, err := os.Open("progress.json")
	 if err != nil {
		return err.Error()
	   }
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
 if err != nil {
		return err.Error()
	   }
progress.Points += 5
	data, err := os.OpenFile("progress.json", os.O_WRONLY|os.O_CREATE, 0644)
	 if err != nil {
		return err.Error()
	   }
	defer data.Close()
	err = json.NewEncoder(data).Encode(&progress)
 if err != nil {
		return err.Error()
	   }
	   
return "+5"
}

	


}
func Currpoints()int{
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {}
	return progress.Points
}
func CurrFood()int{
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {}
	return progress.Food
}
func CurrWater()int{
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {}
	return progress.Water
}

func Levels()string{
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {}
	points := progress.Points
	if points < 20{
		return "Little Noro"
	}else if points >= 20 && points < 40 {
		// Sounds.Sound("Level.mp3")
		return "Noro"
	}else if points >= 40 && points < 60 {
		// Sounds.Sound("Level.mp3")
		return "Noro-Hero"
	} else if points >= 60 && points < 130 {
		// Sounds.Sound("Level.mp3")
		return "Noronuiem"
	}else if points > 9999 {
		// Sounds.Sound("Level.mp3")
		// Sounds.Sound("Level.mp3")
		
		return "Cheater"
	}


	return "Unknown"

}

func levelsound(){
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {}
	points := progress.Points
	if points < 20{
		
	}else if points == 20 {
		Sounds.Sound("Level.mp3")
	
	}else if points == 40{
		Sounds.Sound("Level.mp3")
	
	} else if points == 60 {
		Sounds.Sound("Level.mp3")
	
	}else if points == 130{
		Sounds.Sound("Level.mp3")
	} else if points > 9999 {
		Sounds.Sound("Level.mp3")
		Sounds.Sound("Level.mp3")
		Sounds.Sound("Level.mp3")
	}
}


	func FeedSpammer()int{
	var progress Progress
	data2, err := os.Open("progress.json")
	 if err != nil {
		// return err.Error()
	   }
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
 if err != nil {
		// return err.Error()
	   }
	   progress.Food = progress.Food - 1
	   progress.Count = progress.Count + 1
	   
	   if progress.Count >= 10 {
		levelsound()
		progress.Count = 0
		progress.Points++
		Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
	   }
	   

	data2, err = os.OpenFile("progress.json", os.O_WRONLY|os.O_CREATE, 0644)
	 if err != nil {
		// return err.Error()/
	   }
	defer data2.Close()
	err = json.NewEncoder(data2).Encode(&progress)
 if err != nil {
		// return err.Error()
	   }
	   
return progress.Count
		

	}
	
		func DrinkSpammer()int{
	var progress Progress
	data2, err := os.Open("progress.json")
	 if err != nil {
		// return err.Error()
	   }
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
 if err != nil {
		// return err.Error()
	   }
	   progress.Water = progress.Water - 1
	   progress.Count = progress.Count + 1
	   
	   if progress.Count >= 10 {
		levelsound()
		progress.Count = 0
		progress.Points++
		Sounds.Sound("roblox-eating-nom-nom-nom.mp3")
	   }
	   

	data2, err = os.OpenFile("progress.json", os.O_WRONLY|os.O_CREATE, 0644)
	 if err != nil {
		// return err.Error()/
	   }
	defer data2.Close()
	err = json.NewEncoder(data2).Encode(&progress)
 if err != nil {
		// return err.Error()
	   }
	   
return progress.Count
		

	}
	