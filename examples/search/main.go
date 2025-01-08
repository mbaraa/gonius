package main

import (
	"encoding/json"
	"fmt"

	"github.com/mbaraa/gonius"
)

func main() {
	client := gonius.NewClient("top-secret-shit")
	results, err := client.Search.Get("lana del rey jealous girl")
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		jsonn, _ := json.MarshalIndent(result, "", "\t")
		fmt.Println("search result", string(jsonn))
	}
}
