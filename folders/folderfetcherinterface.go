package folders

import (
	"github.com/gofrs/uuid"
)

// FolderFetcher defines an interface for fetching folders.
type FolderFetcher interface {
	FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error)
	FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID, pageToken string) ([]*Folder, error)
}
