package util

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

func Test_NewDirectoryHandler(t *testing.T) {
	obj := NewDirectoryHandler()

	assert.NotNil(t, obj)
}

func Test_GetCurrentFirectoryPath(t *testing.T) {
	obj := NewDirectoryHandler()
	path := obj.GetCurrentDirectoryPath()

	assert.NotEmpty(t, path)
	fmt.Printf(path)
}

func Test_GetFileListInDirectory(t *testing.T) {
	obj := NewDirectoryHandler()
	path := obj.GetCurrentDirectoryPath()

	files := obj.GetFileListInDirectory(path)
	//assert.EqualValues(t, 0, len(files))

	if len(files) > 0 {
		fmt.Println(files)
	} else {
		fmt.Println("no files")
	}
}

func Test_MakeDirectory(t *testing.T) {
	obj := NewDirectoryHandler()

	path := "/Users/kazufumiwatanabe/go/src/PixelTool"
	assert.True(t, obj.MakeDirectory(path, "hoge"))
}
