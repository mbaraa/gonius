package gonius

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Client is the Genius client that handles all the different API calls to api.genius.com
type Client struct {
	Account     *AccountService
	Annotations *AnnotationsService
	Artists     *ArtistsService
	Lyrics      *LyricsService
	Search      *SearchService
	Songs       *SongsService
}

func (c *Client) SetPageSize() {
	// per_page=20
}

// NewClient initializes the genius [Client] with the given access token to interact with different api.genius.com calls.
func NewClient(accessToken string) *Client {
	baseGeniusUrl := "https://api.genius.com/"

	lyricsService := &LyricsService{}

	c := &Client{}
	c.Account = &AccountService{}
	c.Annotations = &AnnotationsService{
		gClient: newApiClient(http.MethodGet, baseGeniusUrl+"annotations/", accessToken, nil),
	}
	c.Artists = &ArtistsService{
		gClient: newApiClient(http.MethodGet, baseGeniusUrl+"artists/", accessToken, nil),
	}
	c.Lyrics = lyricsService
	c.Search = &SearchService{
		gClient: newApiClient(http.MethodGet, baseGeniusUrl+"search/", accessToken, nil),
	}
	c.Songs = &SongsService{
		gClient: newApiClient(http.MethodGet, baseGeniusUrl+"songs/", accessToken, nil),
		lyrics:  lyricsService,
	}

	return c
}

// ApiResponse is the general response structure that is received from different api calls from api.genius.com
type ApiResponse struct {
	Meta *struct {
		Status int `json:"status,omitempty"`
	} `json:"meta,omitempty"`
	Response *struct {
		Annotation *Annotation `json:"annotation,omitempty"`
		Song       *Song       `json:"song,omitempty"`
		Songs      []Song      `json:"songs,omitempty"`
		Artist     *Artist     `json:"artist,omitempty"`
		Hits       []Hit       `json:"hits,omitempty"`
	} `json:"response,omitempty"`
}

type apiClient struct {
	client      *http.Client
	req         *http.Request
	initialPath string
}

func newApiClient(method, requestPath, token string, body io.Reader) *apiClient {
	req, err := http.NewRequest(method, requestPath, body)
	if err != nil {
		return nil
	}

	a := &apiClient{
		client:      &http.Client{},
		req:         req,
		initialPath: requestPath,
	}
	a.setHeader("Authorization", "Bearer "+token)
	// setting response text_format to plain, so it's readable by the application,
	// rather than using dom or html which need furthor pasring.
	a.setQueryParam("text_format", "plain")

	return a
}

func (a *apiClient) setPath(path string) error {
	if _, err := url.Parse(path); err != nil {
		return err
	}
	a.req.URL.Path = path

	return nil
}

func (a *apiClient) appendToPath(path string) error {
	a.req.URL.Path += path
	return nil
}

func (a *apiClient) setQueryParam(key, value string) error {
	q := a.req.URL.Query()
	q.Set(key, value)
	a.req.URL.RawQuery = q.Encode()

	return nil
}

func (a *apiClient) setHeader(key, value string) error {
	a.req.Header.Set(key, value)
	return nil
}

func (a *apiClient) callEndpoint() (ApiResponse, error) {
	defer a.reset()

	var res ApiResponse
	resp, err := a.client.Do(a.req)
	if err != nil {
		return ApiResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return ApiResponse{}, ErrInvalidToken
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return ApiResponse{}, err
	}

	switch res.Meta.Status {
	case http.StatusOK:
		return res, nil
	case http.StatusNotFound:
		return ApiResponse{}, ErrNotFound
	default:
		return ApiResponse{}, ErrApiError
	}
}

func (a *apiClient) reset() {
	a.req.URL.Path = a.initialPath
	a.setQueryParam("text_format", "plain")
}
