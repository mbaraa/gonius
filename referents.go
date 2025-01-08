package gonius

type ReferentsService struct {
	gClient *apiClient
}

type Referent struct {
	Type           string `json:"_type,omitempty"`
	AnnotatorId    int    `json:"annotator_id,omitempty"`
	AnnotatorLogin string `json:"annotator_login,omitempty"`
	ApiPath        string `json:"api_path,omitempty"`
	Classification string `json:"classification,omitempty"`
	Featured       bool   `json:"featured,omitempty"`
	Fragment       string `json:"fragment,omitempty"`
	Id             int    `json:"id,omitempty"`
	IsDescription  bool   `json:"is_description,omitempty"`
	Path           string `json:"path,omitempty"`
	Range          *struct {
		Start       string `json:"start,omitempty"`
		StartOffset string `json:"startOffset,omitempty"`
		End         string `json:"end,omitempty"`
		EndOffset   string `json:"endOffset,omitempty"`
		Before      string `json:"before,omitempty"`
		After       string `json:"after,omitempty"`
		Content     string `json:"content,omitempty"`
	} `json:"range,omitempty"`
	SongId               int    `json:"song_id,omitempty"`
	URL                  string `json:"url,omitempty"`
	VerifiedAnnotatorIds []int  `json:"verified_annotator_ids,omitempty"`
	Annotatable          *struct {
		ApiPath   string `json:"api_path,omitempty"`
		Context   string `json:"context,omitempty"`
		Id        int    `json:"id,omitempty"`
		ImageURL  string `json:"image_url,omitempty"`
		LinkTitle string `json:"link_title,omitempty"`
		Title     string `json:"title,omitempty"`
		Type      string `json:"type,omitempty"`
		URL       string `json:"url,omitempty"`
	} `json:"annotatable,omitempty"`
	Annotations []Annotation `json:"annotations,omitempty"`
}

type ReferentsParams struct {
	WebPageId   int    `url:"web_page_id,omitempty"`
	CreatedById int    `url:"create_by_id,omitempty"`
	SongId      int    `url:"song_id,omitempty"`
	TextFormat  string `url:"text_format,omitempty"`
}

func (s *ReferentsService) GetBySongId(id int) ([]Referent, error) {
	var err error
	// params := &ReferentsParams{SongID: ID, TextFormat: "plain"}
	res := new(ApiResponse)
	// s.client.base.Get("referents").QueryStruct(params).Receive(res, err)
	if err != nil {
		return nil, err
	}
	return res.Response.Referents, nil
}

func (s *ReferentsService) GetByUserId(id int) ([]Referent, error) {
	var err error
	// params := &ReferentsParams{CreatedByID: ID, TextFormat: "plain"}
	res := new(ApiResponse)
	// s.client.base.Get("referents").QueryStruct(params).Receive(res, err)
	if err != nil {
		return nil, err
	}
	return res.Response.Referents, nil
}

func (s *ReferentsService) GetByWebPageId(id int) ([]Referent, error) {
	var err error
	// params := &ReferentsParams{WebPageID: ID, TextFormat: "plain"}
	res := new(ApiResponse)
	// s.client.base.Get("referents").QueryStruct(params).Receive(res, err)
	if err != nil {
		return nil, err
	}
	return res.Response.Referents, nil
}
