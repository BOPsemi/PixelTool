/*
Definition of white pixel
*/

package models

import (
	"PixelTool/util"
	"strconv"
)

/*
WhitePixel :white pixel structure
*/
type WhitePixel struct {
	Level int // level, unit is DN
	Count int // count, not ppm
}

/*
WhitePixelMapper : white pixel mapper
*/
func WhitePixelMapper(data []string) (*WhitePixel, bool) {
	wp := new(WhitePixel)
	status := false

	// mapping
	if len(data) > 0 {
		/*
			strToInt :convert string to Int
			If the error was detected, the function return -1
		*/
		strToInt := func(str string) int {
			number, err := strconv.Atoi(str)
			if err != nil {
				number = -1
			}
			return number
		}

		/*
			mapping to wp structure
		*/
		wp.Level = strToInt(data[0])
		wp.Count = strToInt(data[1])

		// status update
		status = true

	}

	return wp, status
}

/*
ReadWhitePixel :read white pixel CSV file and map the data to object
*/
func ReadWhitePixel(path string) []WhitePixel {
	// initialize buffer
	wps := make([]WhitePixel, 0)

	// setup csv reader
	reader := util.NewIOUtil()

	// read csv file
	rawdata, status := reader.ReadCSVFile(path)

	// read csv was successful
	if status {
		if len(rawdata) > 0 {
			for _, data := range rawdata {
				wp, mappingstatus := WhitePixelMapper(data)
				if mappingstatus {
					wps = append(wps, *wp)
				}
			}
		}
	}
	return wps
}
