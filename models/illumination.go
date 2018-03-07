package models

import (
	"PixelTool/util"
	"strconv"
)

/*
Illumination :illumination structure
*/
type Illumination struct {
	Wavelength int
	Intensity  float64
}

/*
IlluminationMapper : data mapper
*/
func IlluminationMapper(data []string) (*Illumination, bool) {
	ill := new(Illumination)
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
		ill.Wavelength = strToInt(data[0])
		ill.Intensity = strToFloat64(data[1])

		// status update
		status = true

	}

	return ill, status
}

/*
ReadIllumination :read Illumination CSV file and map the data to object
*/
func ReadIllumination(path string) []Illumination {
	ills := make([]Illumination, 0)

	// setup csv reader
	reader := util.NewIOUtil()

	// read csv file
	rawdata, status := reader.ReadCSVFile(path)

	// read csv was successful
	if status {
		if len(rawdata) > 0 {
			for _, data := range rawdata {
				ill, mappingstatus := IlluminationMapper(data)
				if mappingstatus {
					ills = append(ills, *ill)
				}
			}
		}
	}
	return ills
}
