package folders

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
)

func GetAllFolders(ff FolderFetcher, req *FetchFolderRequest) (*FetchFolderResponse, error) {
	/*
		Summary added by Bharat Naganath:
		The function GetAllFolder is created to retrieve all folders associated with
		an organisation ID provided in the request variable "req".

		This function takes a single parameter, and it is a pointer to the
		struct FetchFolderRequest. It gives two responses, the first one being the
		pointer to FetchFolderResponse and the second one is error.

		Usage:
				req := &folders.FetchFolderRequest{
					OrgID: uuid.FromStringOrNil("abd-efh-ijk-lmo-pst"),
				}
				res, err := folders.GetAllFolders(req)

	*/

	// The variables f1 and fs are not being used, therefore they have been commented out.
	//var (
	//	err error
	//	f1  Folder
	//	fs  []*Folder
	//)

	f, err := ff.FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	// This part of code is commented out, because the "for loop" seems unnecessary.
	//f := []Folder{}
	//for k, v := range r {
	//	f = append(f, *v)
	//}

	var fp []*Folder
	for _, v1 := range f {
		fp = append(fp, v1)
	}

	//var ffr *FetchFolderResponse
	//ffr = &FetchFolderResponse{Folders: fp}
	ffr := &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// FetchAllFoldersByOrgID Added named return parameters "resFolder" and "err", as part of enhancement of adding recover functions.
func FetchAllFoldersByOrgID(orgID uuid.UUID) (resFolder []*Folder, err error) {

	// Adding recover function to avoid program termination on panic.
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("panic occurred: %s", r))
		}
	}()

	folders := GetSampleData()

	resFolder = []*Folder{}
	for _, folder := range folders {
		// Do not return deleted values
		if folder.OrgId == orgID && !folder.Deleted {
			resFolder = append(resFolder, folder)
		}
	}

	// Originally the error was returned as "nil"
	//return resFolder, nil
	// This needs to be corrected, instead return the actual error.

	return resFolder, err
}
