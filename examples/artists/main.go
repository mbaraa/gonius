package main

import (
	"encoding/json"
	"fmt"

	"github.com/mbaraa/gonius"
)

func main() {
	client := gonius.NewClient("top-secret-shit")
	artist, err := client.Artists.Get("15740")
	if err != nil {
		panic(err)
	}

	artistJson, err := json.MarshalIndent(artist, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("artist: %s\n", artistJson)

	artistSongs, err := client.Artists.GetSongs("15740", gonius.ArtistSongsSortPopularity)
	if err != nil {
		panic(err)
	}

	artistSongsJson, err := json.MarshalIndent(artistSongs, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("artist songs: %s\n", artistSongsJson)
}
