/*
Definition of color code
Basically, this structure handles Macbeth color code
*/

package models

import (
	"PixelTool/util"
	"strconv"
)

/*
ColorCode :Macbeth Color Code structure
*/
type ColorCode struct {
	Number int    // code number
	Name   string // code name
	R      uint8  // red, should be unit8
	G      uint8  // green, should be unit8
	B      uint8  // blue, should be unit8
	A      uint8  // a, should be unit8
}

/*
ColorCodeMapper :mapper for ColorCode
*/
func ColorCodeMapper(data []string) (*ColorCode, bool) {
	// initialize status
	status := false

	// initialize ColorCode Object
	code := new(ColorCode)

	// mapping
	if len(data) != 0 {

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
			strToUint8 :convert string to uint8
			If the error was detected, the function return 0
		*/
		strToUint8 := func(str string) uint8 {
			number, err := strconv.Atoi(str)
			if err != nil {
				number = 0
			}

			return uint8(number)
		}

		/*
			Mapping
		*/
		code.Number = strToInt(data[0])
		code.Name = data[1]
		code.R = strToUint8(data[2])
		code.G = strToUint8(data[3])
		code.B = strToUint8(data[4])
		code.A = 255

		// update status
		status = true
	}

	return code, status
}

/*
ReadColorCode :read color code CSV file and map the data to object
*/
func ReadColorCode(path string) []ColorCode {
	// initialize buffer
	colorcodes := make([]ColorCode, 0)

	// setup csv reader
	reader := util.NewIOUtil()

	// read csv file
	rawdata, status := reader.ReadCSVFile(path)

	// read csv was successful
	if status {
		if len(rawdata) > 0 {
			for _, data := range rawdata {
				colorcode, mappingstatus := ColorCodeMapper(data)
				if mappingstatus {
					colorcodes = append(colorcodes, *colorcode)
				}
			}
		}
	}

	return colorcodes
}
