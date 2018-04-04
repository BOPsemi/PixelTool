package controllers

import (
	"PixelTool/models"
)

/*
This controller defines spectrum response
 in ; device QE, sustrate
 out ; spectrum response
*/

/*
ResponseController : response controller
*/
type ResponseController interface {
	ReadResponseData(filepath map[string]string) bool
	CalculateResponse(ill models.IlluminationCode, startWave, stopWave, step int, normPatchNumber int) bool
}

// definition of structure
type responseController struct {
	deviceQE     []models.DeviceQE       // Deivce QE data stocker
	colorChecker [][]models.ColorChecker // Color Checker data stocker
	d65          []models.Illumination   // D65 illumination data stocker
	illA         []models.Illumination   // Illumination-A data stocker

	// map data
	deviceQEGrMap   map[int]float64 // map
	deviceQEGbMap   map[int]float64 // map
	deviceQERedMap  map[int]float64 // map
	deviceQEBlueMap map[int]float64 // map

	colorCheckerMap []map[int]float64
	d65Map          map[int]float64
	illAMap         map[int]float64
}

/*
NewResponseController :initializer of response controller
*/
func NewResponseController() ResponseController {
	obj := new(responseController)

	// initialie properties
	obj.deviceQE = make([]models.DeviceQE, 0)
	obj.colorChecker = make([][]models.ColorChecker, 0)
	obj.d65 = make([]models.Illumination, 0)
	obj.illA = make([]models.Illumination, 0)

	return obj
}

/*
ReadResponseData
 - in	;filepath map[string]string
 - out	;bool
*/
func (rc *responseController) ReadResponseData(filepath map[string]string) bool {
	status := false

	// deviceQE
	rc.deviceQE = models.ReadDeviceQE(filepath["DeviceQE"])

	// color checker
	rc.colorChecker = models.ReadColorChecker(filepath["ColorChecker"])

	// D65 Illumination
	rc.d65 = models.ReadIllumination(filepath["D65"])

	// A Illumination
	rc.illA = models.ReadIllumination(filepath["IllA"])

	// check result
	if len(rc.deviceQE)*len(rc.colorChecker)*len(rc.d65)*len(rc.illA) != 0 {
		// waiting list

		/*
			TODO : need go routine to suppress calculation time
		*/

		// device QE mapping
		gr := make(map[int]float64, 0)
		gb := make(map[int]float64, 0)
		r := make(map[int]float64, 0)
		b := make(map[int]float64, 0)

		for _, data := range rc.deviceQE {
			wavelength := data.GetWavelength()
			gr[wavelength] = data.GetGrSignal()
			gb[wavelength] = data.GetGbSignal()
			r[wavelength] = data.GetRedSignal()
			b[wavelength] = data.GetBlueSignal()
		}

		// D65 mapping
		d65 := make(map[int]float64, 0)
		for _, data := range rc.d65 {
			d65[data.GetWavelangth()] = data.GetIntensity()
		}

		// illA mapping
		illA := make(map[int]float64, 0)
		for _, data := range rc.illA {
			illA[data.GetWavelangth()] = data.GetIntensity()
		}

		// color checker
		colorChecker := make([]map[int]float64, 0)
		for _, obj := range rc.colorChecker {
			checker := make(map[int]float64, 0)
			for _, data := range obj {
				checker[data.GetWavelength()] = data.GetIntensity()
				colorChecker = append(colorChecker, checker)
			}
		}

		// upload all data
		rc.deviceQEGrMap = gr
		rc.deviceQEGbMap = gb
		rc.deviceQERedMap = r
		rc.deviceQEBlueMap = b

		rc.illAMap = illA
		rc.d65Map = d65

		rc.colorCheckerMap = colorChecker

		status = true
	}

	return status
}

/*
CalculateResponse
 - in	;ill models.IlluminationCode, startWave, stopWave, step int, normPatchNumber int
 - out	;bool
*/
func (rc *responseController) CalculateResponse(ill models.IlluminationCode, startWave, stopWave, step int, normPatchNumber int) bool {
	status := false

	return status
}
