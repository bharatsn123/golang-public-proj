package folders

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

const offsetValue = 5

func GetAllFoldersPaginated(ff FolderFetcher, req *FetchFolderRequest) (*FetchFolderResponse, error) {
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

	f, err := ff.FetchAllFoldersByOrgIDPaginated(req.OrgID, req.PageToken)
	if err != nil {
		return nil, err
	}

	var fp []*Folder
	for _, v1 := range f {
		fp = append(fp, v1)
	}
	nextTokenNumber, err := GetNextTokenNumber()
	ffr := &FetchFolderResponse{Folders: fp, NextToken: nextTokenNumber}
	return ffr, nil
}

// FetchAllFoldersByOrgID Added named return parameters "resFolder" and "err", as part of enhancement of adding recover functions.
func FetchAllFoldersByOrgIDPaginated(orgID uuid.UUID, pageToken string) (resFolder []*Folder, err error) {

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
	res, err := PaginateExistingData(resFolder, pageToken)

	return res, err
}

/*
PaginateExistingData
The Pagination Logic:

	Here a new JSON file called pagination.json is created locally,
	This is where the next token is stored, which is mapped to next index
	of the current list of folders.
		Example contents of Pagination.json: {"offsetId":"16","token":"2b608f8aab32"}

	This token is used in the subsequent calls to fetch the next 5 records.

	Assumptions:
	1) GetSampleData() returns the same entries everytime. Ideally, the results are
		stored in database instead of sample.json file. And if that's the case, it
		would allow the pagination logic to leverage database and cache the query
		results. The cached data can be iterated, using tokens.
	2) User wants 5 records at once. The logic can be further improved to have a
		dynamic value to the window size.
*/
func PaginateExistingData(resFolder []*Folder, token string) (paginatedFolder []*Folder, err error) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, "pagination.json")
	var currentToken string
	if len(token) > 0 {
		currentToken, err = GetNextTokenNumber()
		if err != nil {
			fmt.Printf("could not retrieve current token number: %s\n", err)
			return nil, err
		}
	}
	if len(token) == 0 {
		// This means that this is the first request, and we need to generate a new token.
		newToken := GenerateSecureToken(6)
		// find the new offset ID
		nextId := offsetValue + 1
		if cap(resFolder) < offsetValue {
			paginatedFolder = resFolder[:]
			newToken = ""
		} else {
			paginatedFolder = resFolder[:offsetValue]
		}
		map1 := map[string]string{
			"token":    newToken,
			"offsetId": strconv.Itoa(nextId),
		}
		jsonString, err := json.Marshal(map1)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return nil, err
		}
		os.WriteFile(filePath, jsonString, os.ModePerm)
	} else if token == currentToken {
		offsetId, err := GetNextOffsetId()
		if err != nil {
			fmt.Printf("could not get next offset id: %s\n", err)
			return nil, err
		}
		i, err := strconv.Atoi(offsetId)
		if err != nil {
			// ... handle error
			panic(err)
		}
		newToken := GenerateSecureToken(6)
		newOffset := i + offsetValue
		if cap(resFolder[i:]) < offsetValue {
			paginatedFolder = resFolder[:]
			newToken = ""
		} else {
			paginatedFolder = resFolder[i:newOffset]
		}
		map1 := map[string]string{
			"token":    newToken,
			"offsetId": strconv.Itoa(newOffset),
		}
		jsonString, err := json.Marshal(map1)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return nil, err
		}
		os.WriteFile(filePath, jsonString, os.ModePerm)

	} else {
		err = errors.New(fmt.Sprintf("Invalid Token Number: %s", currentToken))
		fmt.Printf("error occurred: %s\n", err)
		return nil, err
	}
	return paginatedFolder, err
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func GetNextTokenNumber() (token string, err error) {
	// Read the json file and get the token number
	dataMap, err := GetPaginationData()
	if err != nil {
		return "", err
	}
	return dataMap["token"].(string), nil
}

func GetNextOffsetId() (token string, err error) {
	// Read the json file and get the token number
	dataMap, err := GetPaginationData()
	if err != nil {
		return "", err
	}
	return dataMap["offsetId"].(string), nil
}

func GetPaginationData() (map[string]interface{}, error) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, "pagination.json")
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var dataMap map[string]interface{}
	err = json.Unmarshal(fileData, &dataMap)
	if err != nil {
		return nil, err
	}
	return dataMap, nil
}
