package viewcontrollers

import (
	"PixelTool/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var path map[string]string
var elm []float64
var gamma float64
var refPatchNumber int

func initSetupFile() {
	// initialize file path
	//path := make(map[string]string, 0)
	path["DeviceQE"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/device_QE.csv"
	path["ColorChecker"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/Macbeth_Color_Checker.csv"
	path["D65"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/illumination_D65.csv"
	path["IllA"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/illumination_A.csv"

	// initialize linear matrix elements
	elm = []float64{0.136, 0.061, 0.104, 0.063, 0.041, 0.057}

	// initialize gamma coefficient
	gamma = 0.43

	// initialize reference pathc number
	refPatchNumber = 22

}

func Test_NewDeviceResponseViewController(t *testing.T) {
	obj := NewDeviceResponseViewController()

	assert.NotNil(t, obj)
}

func Test_ReadResponseRawData(t *testing.T) {
	obj := NewDeviceResponseViewController()

	// make file list for testing
	filepath := make(map[string]string, 0)
	filepath["DeviceQE"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/device_QE.csv"
	filepath["ColorChecker"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/Macbeth_Color_Checker.csv"
	filepath["D65"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/illumination_D65.csv"
	filepath["IllA"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/illumination_A.csv"

	// read test
	status := obj.ReadResponseRawData(filepath)
	assert.True(t, status)

	// calculate Device Response
	obj.CalculateDeviceResponse(models.D65, 400, 700, 5, 0.42, 22)

	/*
		// debugging
			rawdata := obj.RawData()
			rawresdata := obj.RawResponseData()
			for index, data := range rawdata {
				fmt.Println(index+1, data)
			}
			fmt.Println("----------")
			for index, data := range rawresdata {
				fmt.Println(index+1, data)
			}
	*/

	// calculate linear matrix
	elm := []float64{0.136, 0.061, 0.104, 0.063, 0.041, 0.057}
	status = obj.CalculateLinearMatrix(elm)
	assert.True(t, status)

	/*
		// debugging
		linearizedResData := obj.LinearizedResponseData()
		for index, data := range linearizedResData {
			fmt.Println(index+1, data)
		}
	*/

	// calculate white balance
	redGain, blueGain := obj.CalculateWhiteBalanceGain(22)

	for index, data := range obj.LinearizedResponseData() {
		red := data[0] * redGain
		green := data[1]
		blue := data[2] * blueGain

		fmt.Println(index+1, red, green, blue)
	}
}
