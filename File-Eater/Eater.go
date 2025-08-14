package Eater

import(
	"os"
	// "log"
	"encoding/json"
	"main/Sounds"
	"time"
	"strings"
)
type Progress struct {
	Points int `json:"points"`
	Count  int `json:"count"`
	Food int `json:"food"`
	Water int `json:"water"`
	Time time.Time `json:"time"`
	Hunger int64 `json:"hunger"`
	Dif time.Duration `json:"dif"`
}

func Currhunger() int64 {
	var progress Progress
	data2, err := os.Open("progress.json")
	if err != nil {
		return 0
	}
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
	if err != nil {
		return 0
	}
		progress.Dif = time.Duration(time.Since(progress.Time).Seconds())
	if progress.Dif >= 20 {
		if progress.Hunger > 0 {
		progress.Hunger = progress.Hunger - int64(progress.Dif / 20) * 3
		progress.Time = time.Now()
		if progress.Hunger < 0{
			progress.Hunger = 0
		}
		
}}
data, err := os.OpenFile("progress.json", os.O_WRONLY|os.O_CREATE, 0644)
	 if err != nil {
		// return err.Error()
	   }
	defer data.Close()
	err = json.NewEncoder(data).Encode(&progress)
 if err != nil {
		// return err.Error()
	   }
	return progress.Hunger

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
	// progress.Dif = time.Duration(time.Since(progress.Time).Seconds())
	   if err != nil{
		return err.Error()
	   }else{
	// if progress.Dif >= 5 {
	// 	if progress.Hunger > 0 {
	// 	progress.Hunger = progress.Hunger - int64(progress.Dif / 5) * 5
	// 	if progress.Hunger < 0{
	// 		progress.Hunger = 0
	// 	}
	// 	progress.Time = time.Now()
	// }}

	if progress.Hunger > 50 {
	if progress.Points >= 60 {
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
	}else if progress.Hunger < 50{
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
		progress.Hunger = progress.Hunger + 5
		if progress.Points > 60{
			progress.Points = progress.Points + 5 
		}else {
			progress.Points = progress.Points + 3
		}
		data, err := os.OpenFile("progress.json", os.O_WRONLY|os.O_CREATE, 0644)
	 if err != nil {
		return err.Error()
	   }
	defer data.Close()
	err = json.NewEncoder(data).Encode(&progress)
 if err != nil {
		return err.Error()
	   }
	   
return "You are still hungery +5"

	}
	   }
	return "1"
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
	} else if points >= 60 && points < 9998 {
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
	 if err != nil {
		// return err.Error()
	   }
	defer data2.Close()
	err = json.NewDecoder(data2).Decode(&progress)
 if err != nil {
		// return err.Error()
	   }
	   points := progress.Points
	if points < 20{
		
	}else if points == 20 {
		Sounds.Sound("Level.mp3")
		progress.Food += 100
		
	
	}else if points == 40{
		Sounds.Sound("Level.mp3")
		
		progress.Water += 100
	
	} else if points == 60 {
		Sounds.Sound("Level.mp3")
		progress.Food += 100
		
	
	}else if points == 130{
		Sounds.Sound("Level.mp3")
		
		progress.Water += 100
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
	  


	   if progress.Food > 0 {
		progress.Count = progress.Count + 1
	  			if progress.Hunger >= 50 && progress.Count >= 10{
				progress.Hunger = progress.Hunger + 5
				if progress.Hunger > 100{
					progress.Hunger = 100
				}
	   progress.Food = progress.Food - 1
	   
			} else if progress.Hunger < 50 && progress.Count >= 10{
				progress.Hunger = progress.Hunger + 8
				progress.Food = progress.Food - 1
				
			}
	   
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
	return 0

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
	  
// 	progress.Dif = time.Duration(time.Since(progress.Time).Seconds())
// 	if progress.Dif >= 20 {
// 		if progress.Hunger > 0 {
// 		progress.Hunger = progress.Hunger - int64(progress.Dif / 20) * 3
// 		if progress.Hunger < 0{
// 			progress.Hunger = 0
// 		}
// 		progress.Time = time.Now()
// }}

	   if progress.Water > 0 {
		progress.Count = progress.Count + 1
	  			if progress.Hunger >= 50 && progress.Count >= 10{
				progress.Hunger = progress.Hunger + 5
				if progress.Hunger > 100{
					progress.Hunger = 100
				}
	   progress.Water = progress.Water - 1
	   
			} else if progress.Hunger < 50 && progress.Count >= 10{
				progress.Hunger = progress.Hunger + 8
				progress.Water = progress.Water - 1
				
			}
	   
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
	return 0

	}
	

	func Currenttime()string{
		curtime := time.Now().Format("03:04 PM")
		return curtime
	}
	func Daynight()string{
		if strings.Contains(Currenttime(),"A"){
			return "You are in the morning"
		}else{
			return "You are at night"
		}

	}