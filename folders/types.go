package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID     uuid.UUID
	PageToken string // Token for the starting point of the current page
}

type FetchFolderResponse struct {
	Folders   []*Folder
	NextToken string // Token for the next page
}
