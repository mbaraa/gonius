package main

import (
	"encoding/json"
	"fmt"

	"github.com/mbaraa/gonius"
)

func main() {
	client := gonius.NewClient("top-secret-shit")
	song, err := client.Songs.Get("3130730")
	if err != nil {
		panic(err)
	}

	fmt.Printf("song.URL: %v\n", song.Description)

	songJson, err := json.MarshalIndent(song, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("song: %s\n", songJson)

}
