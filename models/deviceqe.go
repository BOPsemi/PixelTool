/*
The model of device QE (quantum efficiency)
*/

package models

import (
	"PixelTool/util"
	"strconv"
)

/*
DeviceQE :device QE structure
*/
type DeviceQE struct {
	Wavelength int
	Gr         float64 // Gr channel QE
	Gb         float64 // Gb channel QE
	R          float64 // Red channel QE
	B          float64 // Blue channel QE
}

/*
DeviceQEMapper :data mapper
*/
func DeviceQEMapper(data []string) (*DeviceQE, bool) {
	qe := new(DeviceQE)
	status := false

	// mapping
	if len(data) > 0 {
		/*
			strToFloat64 :converter from string to Float64
		*/
		strToFloat64 := func(str string) float64 {
			number, err := strconv.ParseFloat(str, 64)
			if err != nil {
				number = 0.0
			}

			return number
		}

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
			mapping
		*/
		qe.Wavelength = strToInt(data[0])
		qe.Gr = strToFloat64(data[1])
		qe.Gb = strToFloat64(data[2])
		qe.R = strToFloat64(data[3])
		qe.B = strToFloat64(data[4])

		// status update
		status = true

	}

	return qe, status
}

/*
ReadDeviceQE :read device QE CSV file and map the data to object
*/
func ReadDeviceQE(path string) []DeviceQE {
	qes := make([]DeviceQE, 0)

	// setup csv reader
	reader := util.NewIOUtil()

	// read csv file
	rawdata, status := reader.ReadCSVFile(path)

	// read csv was successful
	if status {
		if len(rawdata) > 0 {
			for _, data := range rawdata {
				qe, mappingstatus := DeviceQEMapper(data)
				if mappingstatus {
					qes = append(qes, *qe)
				}
			}
		}
	}
	return qes
}
