package gonius

type SearchService struct {
	gClient *apiClient
}

type Hit struct {
	HighLights any    `json:"highlights,omitempty"`
	Index      string `json:"string,omitempty"`
	Type       string `json:"type,omitempty"`
	Result     *struct {
		AnnotationCount          int    `json:"annotation_count,omitempty"`
		ApiPath                  string `json:"api_path,omitempty"`
		FullTitle                string `json:"full_title,omitempty"`
		HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url,omitempty"`
		Id                       int    `json:"id,omitempty"`
		LyricsOwnerId            int    `json:"lyrics_owner_id,omitempty"`
		LyricsState              string `json:"lyrics_state,omitempty"`
		Path                     string `json:"path,omitempty"`
		PyongsCount              int    `json:"pyongs_count,omitempty"`
		SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url,omitempty"`
		SongArtImageURL          string `json:"song_art_image_url,omitempty"`
		Stats                    *struct {
			UnreviewedAnnotations int  `json:"unreviewed_annotations,omitempty"`
			Concurrents           int  `json:"concurrents,omitempty"`
			Hot                   bool `json:"hot,omitempty"`
			PageViews             int  `json:"page_views,omitempty"`
		} `json:"stats,omitempty"`
		Title             string  `json:"title,omitempty"`
		TitleWithFeatured string  `json:"title_with_featured,omitempty"`
		URL               string  `json:"url,omitempty"`
		PrimaryArtist     *Artist `json:"primary_artist,omitempty"`
	} `json:"result,omitempty"`
}

func (s *SearchService) Get(q string) ([]Hit, error) {
	s.gClient.setQueryParam("q", q)

	res, err := s.gClient.callEndpoint()
	if err != nil {
		return nil, err
	}

	return res.Response.Hits, nil
}
