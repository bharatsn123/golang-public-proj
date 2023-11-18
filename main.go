package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

// RealFolderFetcher is an implementation of the FolderFetcher interface that uses real data.
type RealFolderFetcher struct{}

func main() {
	// Usage with Token:
	//req := &folders.FetchFolderRequest{
	//	OrgID:     uuid.FromStringOrNil(folders.DefaultOrgID),
	//	PageToken: "6de349030601",
	//}

	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	// Create an instance of RealFolderFetcher - Interface mocked for unit tests
	fetcher := folders.RealFolderFetcher{}

	//res, err := folders.GetAllFolders(fetcher, req)
	res, err := folders.GetAllFoldersPaginated(fetcher, req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	folders.PrettyPrint(res)
}
