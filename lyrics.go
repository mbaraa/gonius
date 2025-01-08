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
	sectionsPattern = regexp.MustCompile(`(\[Chorus\]|\[Pre-Chorus\]|\[Outro\]|\[Intro\]|\[Verse \d+\]|\[Bridge\]|\[Produced.*\])`)
	linesPattern    = regexp.MustCompile(`\n{2,}|\r{2,}`)
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

func newLyrics(text []string) Lyrics {
	// sectionsCount := sectionsPattern.FindAllString(strings.Join(text, ""), -1)
	fixedLyrics := make([]string, 0, len(text))
	for _, lyric := range text {
		lyric = strings.TrimSpace(strings.Trim(strings.Trim(lyric, "\r"), "\n"))
		if sectionsPattern.MatchString(lyric) {
			fixedLyrics = append(fixedLyrics, "\n"+lyric)
		} else {
			fixedLyrics = append(fixedLyrics, lyric)
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

	lyricsRaw := make([]string, 0)

	c.OnHTML("[data-lyrics-container=\"true\"]", func(e *colly.HTMLElement) {
		for _, node := range e.DOM.Nodes {
			for child := range node.Descendants() {
				switch child.Type {
				case html.ElementNode:
				default:
					lyricsRaw = append(lyricsRaw, child.Data)
				}
			}
		}
	})

	err := c.Visit(songUrl + "/lyrics?text_format=plain")
	if err != nil {
		return Lyrics{}, err
	}

	if len(lyricsRaw) > 0 &&
		lyricsRaw[0] == "" || lyricsRaw[0] == " " || lyricsRaw[0] == "\n" || lyricsRaw[0] == "\r" {
		lyricsRaw = lyricsRaw[1:]
	}

	return newLyrics(lyricsRaw), nil
}
