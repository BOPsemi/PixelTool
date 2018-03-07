package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
Directrory handler class handls the works
which are related to directory control
*/

/*
DirectoryHandler :interface of directory handler
*/
type DirectoryHandler interface {
	MakeDirectory(path, name string) bool
	GetCurrentDirectoryPath() string
	GetFileListInDirectory(path string) []string
}

// definition of directory handler
type directoryHandler struct {
}

/*
NewDirectoryHandler :initializer of directory handler
*/
func NewDirectoryHandler() DirectoryHandler {
	obj := new(directoryHandler)

	return obj
}

/*
MakeDirectory :make directory
*/
func (di *directoryHandler) MakeDirectory(path, name string) bool {
	status := false

	// check path and name are not blank
	if path != "" && name != "" {
		fullPath := path + "/" + name

		if !di.checkDirectoryPath(fullPath) {
			// the directory has already been available
			if di.deleteDirectory(fullPath) {
				// delete success
				if di.createDirectory(fullPath) {
					// success to create new directory
					status = true
				}
			}
		} else {
			// the directory is not avilable
			if di.createDirectory(fullPath) {
				// directory creation was successed
				status = true
			}
		}
	}

	return status
}

/*
checkDirectoryPath
	available	: false
	nothing		: true
*/
func (di *directoryHandler) checkDirectoryPath(path string) bool {
	status := true

	// check path name
	_, err := os.Stat(path)
	if err == nil {
		status = false
	}

	return status
}

/*
deleteDeirectory
	true	:success
	false	:failed
*/
func (di *directoryHandler) deleteDirectory(path string) bool {
	status := false

	err := os.Remove(path)
	if err == nil {
		status = true
	}

	return status
}

/*
createDirectory
	true	:success
	false	:failed
*/
func (di *directoryHandler) createDirectory(path string) bool {
	status := false
	err := os.Mkdir(path, 0776)
	if err == nil {
		status = true
	}

	return status
}

/*
CurrentFirectoryPath() string
*/
func (di *directoryHandler) GetCurrentDirectoryPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		path = ""
	}
	return path
}

/*
GetFileListInDirectory :retrun file list in the path
*/
func (di *directoryHandler) GetFileListInDirectory(path string) []string {
	list := make([]string, 0)

	if path != "" {
		if !di.checkDirectoryPath(path) {
			// false -> available
			files, err := ioutil.ReadDir(path)
			if err == nil {
				for _, file := range files {
					fileName := path + file.Name()
					list = append(list, fileName)
				}
			}
		}
	}

	return list
}