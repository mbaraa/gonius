package gonius

type SongsService struct {
	gClient *apiClient
	lyrics  *LyricsService
}

type Song struct {
	FeaturedVideo bool `json:"featured_video,omitempty"`

	AnnotationCount int `json:"annotation_count,omitempty"`
	Id              int `json:"id,omitempty"`
	LyricsOwnerId   int `json:"lyrics_owner_id,omitempty"`
	PyongsCount     int `json:"pyongs_count,omitempty"`

	ApiPath                  string `json:"api_path,omitempty"`
	AppleMusicId             string `json:"apple_music_id,omitempty"`
	AppleMusicPlayerURL      string `json:"apple_music_player_url,omitempty"`
	EmbedContent             string `json:"embed_content,omitempty"`
	FullTitle                string `json:"full_title,omitempty"`
	HeaderImageThumbnailURL  string `json:"header_image_thumbnail_url,omitempty"`
	HeaderImageURL           string `json:"header_image_url,omitempty"`
	LyricsState              string `json:"lyrics_state,omitempty"`
	Path                     string `json:"path,omitempty"`
	SongArtImageThumbnailURL string `json:"song_art_image_thumbnail_url,omitempty"`
	SongArtImageURL          string `json:"song_art_image_url,omitempty"`
	Title                    string `json:"title,omitempty"`
	TitleFeatured            string `json:"title_features,omitempty"`
	URL                      string `json:"url,omitempty"`

	Album             *Album             `json:"album,omitempty"`
	FeaturedArtists   []Artist           `json:"featured_artists,omitempty"`
	Media             []Media            `json:"media,omitempty"`
	PrimaryArtist     *Artist            `json:"primary_artist,omitempty"`
	ProducerArtists   []Artist           `json:"producer_artists,omitempty"`
	SongRelationships []SongRelationship `json:"song_relationships,omitempty"`
	WriterArtists     []Artist           `json:"writer_artists,omitempty"`

	Description struct {
		Plain string `json:"plain,omitempty"`
	} `json:"description,omitempty"`
}

type Album struct {
	Id int `json:"id,omitempty"`

	ApiPath     string `json:"api_path,omitempty"`
	CoverArtURL string `json:"cover_art_url,omitempty"`
	FullTitle   string `json:"full_title,omitempty"`
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`

	Artist *Artist `json:"artist,omitempty"`
}

type Media struct {
	Start int `json:"start,omitempty"`

	NativeURI   string `json:"native_uri,omitempty"`
	Provider    string `json:"provider,omitempty"`
	Type        string `json:"type,omitempty"`
	URL         string `json:"url,omitempty"`
	Attribution string `json:"attribution,omitempty"`
}

type SongRelationship struct {
	Type  string `json:"type,omitempty"`
	Songs []Song `json:"songs,omitempty"`
}

func (s *SongsService) Get(id string) (Song, error) {
	s.gClient.appendToPath(id)

	res, err := s.gClient.callEndpoint()
	if err != nil {
		return Song{}, err
	}

	return *res.Response.Song, nil
}

func (s *SongsService) FetchLyrics(song Song) (Lyrics, error) {
	return s.lyrics.FindForSong(song.URL)
}
