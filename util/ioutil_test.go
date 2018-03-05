package util

import (
	"PixelTool/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewIOUtil(t *testing.T) {
	obj := NewIOUtil()
	assert.NotNil(t, obj)
}

func Test_ReadCSVFile(t *testing.T) {
	path := "/Users/kazufumiwatanabe/go/src/PixelTool/data/macbeth_patch_code.csv"

	// check not fail
	obj := NewIOUtil()
	data, status := obj.ReadCSVFile(path)
	assert.True(t, status)
	assert.NotEmpty(t, data)

	// check file name empty
	data, status = obj.ReadCSVFile("")
	assert.False(t, status)
	assert.Empty(t, data)

	// check file open fail
	data, status = obj.ReadCSVFile("hoge")
	assert.False(t, status)
	assert.Empty(t, data)
}

func Test_ReadCSVFileMapper(t *testing.T) {

	path := "/Users/kazufumiwatanabe/go/src/PixelTool/data/macbeth_patch_code.csv"

	// open data
	obj := NewIOUtil()
	data, _ := obj.ReadCSVFile(path)

	// mapping the data to colorcode str
	codes := make([]models.ColorCode, 0)
	for _, rawdata := range data {
		code, _ := models.ColorCodeMapper(rawdata)
		codes = append(codes, *code)
	}

	fmt.Println(codes)
}
