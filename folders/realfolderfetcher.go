package folders

import "github.com/gofrs/uuid"

// RealFolderFetcher is an implementation of the FolderFetcher interface that uses real data.
type RealFolderFetcher struct{}

// FetchAllFoldersByOrgID is the actual implementation of the method.
func (rf RealFolderFetcher) FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// The actual implementation goes here.
	return FetchAllFoldersByOrgID(orgID)
}

func (rf RealFolderFetcher) FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID, pageToken string) ([]*Folder, error) {
	// The actual implementation goes here.
	return FetchAllFoldersByOrgIDPaginated(orgID, pageToken)
}
