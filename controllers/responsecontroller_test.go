package controllers

import "testing"
import "github.com/stretchr/testify/assert"

func Test_NewResponseController(t *testing.T) {
	obj := NewResponseController()

	assert.NotNil(t, obj)
}

func Test_ReadResponseData(t *testing.T) {
	obj := NewResponseController()

	// make file list for testing
	path := make(map[string]string, 0)
	path["DeviceQE"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/device_QE.csv"
	path["ColorChecker"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/Macbeth_Color_Checker.csv"
	path["D65"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/illumination_D65.csv"
	path["IllA"] = "/Users/kazufumiwatanabe/go/src/PixelTool/data/illumination_A.csv"

	status := obj.ReadResponseData(path)

	assert.True(t, status)

}
