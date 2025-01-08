# gonius

**gonius** is a Go client to access the [Genius API](https://docs.genius.com/).

[![Go Report Card](https://goreportcard.com/badge/github.com/mbaraa/gonius)](https://goreportcard.com/report/github.com/mbaraa/gonius)
[![GoDoc](https://godoc.org/github.com/mbaraa/gonius?status.png)](https://godoc.org/github.com/mbaraa/gonius)

# Contributing

IDK, it would be really nice of you to contribute, check the poorly written [CONTRIBUTING.md](/CONTRIBUTING.md) for more info.

# Roadmap

- [x] Search
- [x] Get artist
- [x] Get artist's songs
- [x] Get annotation
- [x] Get song
- [x] Lyrics
- [ ] Pagination
- [ ] Account
- [ ] Find missing shit using [genius-lyrics](https://www.npmjs.com/package/genius-lyrics) as a reference
- [ ] Lyrics timing?

# Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/mbaraa/gonius"
)

func main() {
	client := gonius.NewClient("top-secret-token-woo-scary")
	results, err := client.Search.Get("lana del rey jealous girl")
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		jsonn, _ := json.MarshalIndent(result, "", "\t")
		fmt.Println("search result", string(jsonn))
	}
}
```

# Support :)

Give a ⭐️ if this project helped you!

---

Made with 🧉 by [Baraa Al-Masri](https://mbaraa.com)
