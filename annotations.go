package gonius

type AnnotationsService struct {
	gClient *apiClient
}

type Annotation struct {
	ApiPath             string `json:"api_path,omitempty"`
	Body                any    `json:"body,omitempty"`
	CommentCount        int    `json:"comment_count,omitempty"`
	Community           bool   `json:"community,omitempty"`
	CustomPreview       bool   `json:"custom_preview,omitempty"`
	HasVoters           bool   `json:"has_voters,omitempty"`
	Id                  int    `json:"id,omitempty"`
	Pinned              bool   `json:"pinned,omitempty"`
	ShareURL            string `json:"share_url,omitempty"`
	Source              string `json:"source,omitempty"`
	State               string `json:"state,omitempty"`
	URL                 string `json:"url,omitempty"`
	Verified            bool   `json:"verified,omitempty"`
	VotesTotal          int    `json:"votes_total,omitempty"`
	CurrentUserMetadata *struct {
		Permissions         []string `json:"permissions,omitempty"`
		ExcludedPermissions []string `json:"excluded_permissions,omitempty"`
		Interactions        *struct {
			Following bool `json:"following,omitempty"`
		} `json:"interactions,omitempty"`
	} `json:"current_user_metadata,omitempty"`
	Authors          []Author `json:"authors,omitempty"`
	CosignedBy       []any    `json:"cosigned_by,omitempty"`
	RejectionComment string   `json:"rejection_comment,omitempty"`
	VerifiedBy       *User    `json:"verified_by,omitempty"`
}

type Author struct {
	Attribution float64 `json:"attribution,omitempty"`
	PinnedRole  string  `json:"pinned_role,omitempty"`
	User        *User   `json:"user,omitempty"`
}

func (s *AnnotationsService) Get(id string) (Annotation, error) {
	s.gClient.appendToPath(id)

	res, err := s.gClient.callEndpoint()
	if err != nil {
		return Annotation{}, err
	}

	return *res.Response.Annotation, nil
}
