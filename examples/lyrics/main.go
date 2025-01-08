package main

import (
	"fmt"

	"github.com/mbaraa/gonius"
)

func main() {
	client := gonius.NewClient("top-secret-shit")
	lyrics, err := client.Lyrics.FindForSong("https://genius.com/Adele-set-fire-to-the-rain-lyrics")
	if err != nil {
		panic(err)
	}

	fmt.Println(lyrics.String())
}
