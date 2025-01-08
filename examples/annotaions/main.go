package main

import (
	"encoding/json"
	"fmt"

	"github.com/mbaraa/gonius"
)

func main() {
	client := gonius.NewClient("top-secret-shit")
	annotation, err := client.Songs.Get("10225840")
	if err != nil {
		panic(err)
	}

	annotationJson, err := json.MarshalIndent(annotation, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("annotation: %s\n", annotationJson)
}
