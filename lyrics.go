package gonius

import (
	"crypto/tls"
	"net/http"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

var (
	// sectionsPattern used to identify special parts of the lyrics.
	sectionsPattern = regexp.MustCompile(`(\[.*\])`)
)

// LyricsService fetches lyrics of a song.
type LyricsService struct {
	gClient *apiClient
}

// Lyrics holds the actual plain text value of the lyrics.
type Lyrics struct {
	parts []string
}

func (l *Lyrics) String() string {
	return strings.TrimSpace(strings.Join(l.parts, "\n"))
}

func (l *Lyrics) Parts() []string {
	return l.parts
}

func newLyrics(text string) Lyrics {
	lyricsParts := strings.Split(text, "\n")
	fixedLyrics := make([]string, 0, len(lyricsParts))
	for _, part := range lyricsParts {
		if part == "" {
			continue
		}
		if sectionsPattern.MatchString(part) {
			fixedLyrics = append(fixedLyrics, "\n"+part)
		} else {
			fixedLyrics = append(fixedLyrics, part)
		}
	}

	return Lyrics{
		parts: fixedLyrics,
	}
}

func (l *LyricsService) FindForSong(songUrl string) (Lyrics, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("genius.com", "api.genius.com"),
		colly.AllowURLRevisit(),
	)
	noSSL := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c.WithTransport(noSSL)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1})

	lyricsRaw := new(strings.Builder)

	c.OnHTML("[data-lyrics-container=\"true\"]", func(e *colly.HTMLElement) {
		for _, node := range e.DOM.Nodes {
			for child := range node.Descendants() {
				switch child.Type {
				case html.ElementNode:
					if child.Data == "br" {
						lyricsRaw.WriteRune('\n')
					}
				default:
					lyricsRaw.WriteString(child.Data)
				}
			}
		}
	})

	err := c.Visit(songUrl + "/lyrics?text_format=plain")
	if err != nil {
		return Lyrics{}, err
	}

	return newLyrics(lyricsRaw.String()), nil
}
