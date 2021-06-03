package main

import (
	"fmt"
	"github.com/dingjiefeng/go-music-reg/v0.1/audio"
)

func main() {
	a := audio.Audio{}
	a.SetFile("F:\\music\\demo.wav")
	a.ResloveContent()
	fmt.Println(a.Header, a.Format, a.Content.ID)
}
