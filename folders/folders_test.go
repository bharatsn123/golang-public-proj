package folders_test

import (
	"errors"
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockFolderFetcher is a mock implementation of FolderFetcher.
type MockFolderFetcher struct{}

func (m MockFolderFetcher) FetchAllFoldersByOrgID(orgID uuid.UUID) (folder []*folders.Folder, err error) {
	// Mock behavior based on orgID
	// Return error for specific orgID, or a list of folders for others
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("panic occurred: %s", r))
		}
	}()

	if orgID.String() == "123e4567-e89b-12d3-a456-426655440000" {
		panic("error opening json")
	} else if orgID.String() == "123e6789-e89b-12d3-a456-426655440000" {
		ret_folder := make([]*folders.Folder, 2, 2)
		return ret_folder, err
	}

	return folder, err
}

func TestGetAllFolders(t *testing.T) {
	mockFetcher := MockFolderFetcher{}

	// Define test cases
	testCases := []struct {
		name          string
		orgID         uuid.UUID
		expectedError bool
		expectedCount int
	}{
		{"WithError", uuid.Must(uuid.FromString("123e4567-e89b-12d3-a456-426655440000")), true, 0},
		{"WithNoFolders", uuid.Must(uuid.NewV4()), false, 0},
		{"WithFolders", uuid.Must(uuid.FromString("123e6789-e89b-12d3-a456-426655440000")), false, 2}, // Assume 2 folders returned for this orgID
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := &folders.FetchFolderRequest{OrgID: tc.orgID}
			resp, err := folders.GetAllFolders(mockFetcher, req)

			if tc.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Len(t, resp.Folders, tc.expectedCount)
			}
		})
	}
}
