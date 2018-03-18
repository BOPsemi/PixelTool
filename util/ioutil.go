/*
This tool provides IO interface
	Read CSV
*/

package util

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

/*
IOUtil :interface of IOUtil object
*/
type IOUtil interface {
	ReadCSVFile(path string) ([][]string, bool)
	StreamOutPNGFile(path, filename string, data *image.RGBA) bool
}

// definition of ioUtil
type ioUtil struct {
	file *os.File // pointer for file
}

/*
NewIOUtil : initializer of IOUtil
*/
func NewIOUtil() IOUtil {
	obj := new(ioUtil)

	return obj
}

/*
ReadCSVFile(path string) [][]string
*/
func (i *ioUtil) ReadCSVFile(path string) ([][]string, bool) {
	// initialize status flag
	status := false

	// initialize buffer
	buffer := make([][]string, 0)

	// check file path is empty or not
	// empty		:false
	// not empty	:true
	if path == "" {
		// in the case of file name is empty
		status = false

	} else {
		// open file
		if i.open(path) {
			// setup reader
			reader := csv.NewReader(i.file)

			// read columns
			for {
				data, err := reader.Read()
				if err == io.EOF {
					break
				}
				buffer = append(buffer, data)
			}

			// status update
			status = true

			// close file
			defer i.file.Close()

		} else {
			// fail to opne file
			defer i.file.Close()
		}
	}

	return buffer, status
}

func (i *ioUtil) open(path string) bool {
	status := false

	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err)
	} else {
		i.file = file
		status = true
	}

	return status
}

/*
StreamOutPNGFile :stream out PNG image to path, need file name
*/
func (i *ioUtil) StreamOutPNGFile(path, filename string, data *image.RGBA) bool {
	status := false

	if (path != "") && (filename != "") && (data != nil) {

		// save png file in the full path
		imageName := path + filename + ".png"

		// file opend
		file, err := os.OpenFile(imageName, os.O_WRONLY|os.O_CREATE, 0600)
		defer file.Close()
		if err == nil {
			// PNG file save in the folder
			png.Encode(file, data)
			status = true
		} else {
			fmt.Println(err)
		}

	}

	return status
}
