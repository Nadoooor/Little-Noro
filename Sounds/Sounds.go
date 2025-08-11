package Sounds

import (
	"github.com/gopxl/beep/v2/mp3"
	"log"
	"os"
	"github.com/gopxl/beep/v2/speaker"
	"time"
)

func Sound(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

}