package gonius

type ArtistsService struct {
	gClient *apiClient
}

type Artist struct {
	IsMemeVerified    bool `json:"is_meme_verified,omitempty"`
	IsVerified        bool `json:"is_verified,omitempty"`
	TranslationArtist bool `json:"translation_artist,omitempty"`

	FollowersCount int `json:"followers_count,omitempty"`
	Id             int `json:"id,omitempty"`
	IQ             int `json:"iq,omitempty"`

	ApiPath        string `json:"api_path,omitempty"`
	HeaderImageURL string `json:"header_image_url,omitempty"`
	ImageURL       string `json:"image_url,omitempty"`
	InstagramName  string `json:"instagram_name,omitempty"`
	Name           string `json:"name,omitempty"`
	TwitterName    string `json:"twitter_name,omitempty"`
	URL            string `json:"url,omitempty"`

	AlternateNames []string `json:"alternate_names,omitempty"`
	User           *User    `json:"user,omitempty"`
}

type User struct {
	ApiPath string `json:"api_path,omitempty"`
	Avatar  *struct {
		Tiny *struct {
			URL         string `json:"url,omitempty"`
			BoundingBox *struct {
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"bounding_box,omitempty"`
		} `json:"tiny,omitempty"`
		Thumb *struct {
			URL         string `json:"url,omitempty"`
			BoundingBox *struct {
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"bounding_box,omitempty"`
		} `json:"thumb,omitempty"`
		Small *struct {
			URL         string `json:"url,omitempty"`
			BoundingBox *struct {
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"bounding_box,omitempty"`
		} `json:"small,omitempty"`
		Medium *struct {
			URL         string `json:"url,omitempty"`
			BoundingBox *struct {
				Width  int `json:"width,omitempty"`
				Height int `json:"height,omitempty"`
			} `json:"bounding_box,omitempty"`
		} `json:"medium,omitempty"`
	} `json:"avatar,omitempty"`
	HeaderImageURL              string `json:"header_image_url,omitempty"`
	HumanReadableRoleForDisplay string `json:"human_readable_role_for_display,omitempty"`
	Id                          int    `json:"id,omitempty"`
	IQ                          int    `json:"iq,omitempty"`
	Login                       string `json:"login,omitempty"`
	Name                        string `json:"string,omitempty"`
	RoleForDisplay              string `json:"role_for_display,omitempty"`
	URL                         string `json:"url,omitempty"`
	CurrentUserMetadata         *struct {
		Permissions         []string `json:"permissions,omitempty"`
		ExcludedPermissions []string `json:"excluded_permissions,omitempty"`
		Interactions        *struct {
			Following bool `json:"following,omitempty"`
		} `json:"interactions,omitempty"`
	} `json:"current_user_metadata,omitempty"`
}

func (s *ArtistsService) Get(id string) (Artist, error) {
	s.gClient.appendToPath(id)

	res, err := s.gClient.callEndpoint()
	if err != nil {
		return Artist{}, err
	}

	return *res.Response.Artist, nil
}

func (s *ArtistsService) GetSongs(artistId string, sort ArtistSongsSort) ([]Song, error) {
	return s.getArtistSongs(getArtistSongsParams{
		artistId: artistId,
		sort:     sort,
	})
}

type ArtistSongsSort string

const (
	ArtistSongsSortTitle      ArtistSongsSort = "title"
	ArtistSongsSortPopularity ArtistSongsSort = "popularity"
)

type getArtistSongsParams struct {
	artistId string
	sort     ArtistSongsSort
}

func (s *ArtistsService) getArtistSongs(params getArtistSongsParams) ([]Song, error) {
	s.gClient.appendToPath(params.artistId)
	s.gClient.appendToPath("/songs")
	s.gClient.setQueryParam("sort", string(params.sort))
	s.gClient.setQueryParam("per_page", "25")

	res, err := s.gClient.callEndpoint()
	if err != nil {
		return nil, err
	}

	return res.Response.Songs, nil
}
