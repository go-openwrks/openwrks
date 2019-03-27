package users

type (
	// ListUsersOptions provides a way of filtering down users.
	ListUsersOptions struct {
		// (optional) Filters users whose Customer Reference starts with this text.
		ReferencePrefix string

		Page  int64
		Limit int64
	}

	// ListUsersResponseMetadata ...
	ListUsersResponseMetadata struct {
		TotalNumberOfRecords int64 `json:"totalNumberOfRecords"`
		TotalNumberOfPages   int64 `json:"totalNumberOfPages"`
		CurrentPageNumber    int64 `json:"pageNumber"`
		PageSize             int64 `json:"pageSize"`
	}

	// ListUsersResponseLinks ...
	ListUsersResponseLinks []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	}

	// ListUsersResponse ...
	ListUsersResponse struct {
		Data     []User                    `json:"data"`
		Metadata ListUsersResponseMetadata `json:"_meta"`
		Links    ListUsersResponseLinks    `json:"_links"`
	}

	// User ...
	User struct {
		UserID    string `json:"userId"`
		Reference string `json:"reference"`
	}
)
