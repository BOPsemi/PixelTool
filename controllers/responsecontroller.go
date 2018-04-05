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
	CalculateResponse(ill models.IlluminationCode, startWave, stopWave, step int, normPatchNumber int) (bool, []models.ChannelResponse)
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

	// ressults
	responses []models.ChannelResponse
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
	obj.responses = make([]models.ChannelResponse, 0)

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
			gr[wavelength] = data.GetGrSignal() * 0.01
			gb[wavelength] = data.GetGbSignal() * 0.01
			r[wavelength] = data.GetRedSignal() * 0.01
			b[wavelength] = data.GetBlueSignal() * 0.01
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
			}
			colorChecker = append(colorChecker, checker)
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
func (rc *responseController) CalculateResponse(ill models.IlluminationCode, startWave, stopWave, step int, normPatchNumber int) (bool, []models.ChannelResponse) {
	status := false
	responses := make([]models.ChannelResponse, 0)

	var illSpectrum map[int]float64
	switch ill {
	case models.D65:
		illSpectrum = rc.d65Map
	case models.IllA:
		illSpectrum = rc.illAMap
	default:
		break
	}

	// calculate each color chart response
	for i := 0; i < 24; i++ {
		// device channel response stocker
		grCh := 0.0
		gbCh := 0.0
		rCh := 0.0
		bCh := 0.0

		// scan wave length
		for wavelength := startWave; wavelength <= stopWave; wavelength += step {
			gr, gb, r, b := rc.calculateEachChannelResponse(
				illSpectrum,
				rc.deviceQEGrMap,
				rc.deviceQEGbMap,
				rc.deviceQERedMap,
				rc.deviceQEBlueMap,
				rc.colorCheckerMap[i],
				wavelength,
			)

			// accumulate response
			grCh += gr
			gbCh += gb
			rCh += r
			bCh += b
		}

		// make channel response object
		res := &models.ChannelResponse{
			CheckerNumber: i + 1,
			Gr:            grCh,
			Gb:            gbCh,
			R:             rCh,
			B:             bCh,
		}

		// stock the chennel response data to stocker
		rc.responses = append(rc.responses, *res)

		// update status
		status = true
	}

	// normarize channel response by ref patch signal
	if status {
		refPatch := rc.responses[normPatchNumber-1]
		refPatchGrGb := (refPatch.Gr + refPatch.Gb) / 2.0

		for _, data := range rc.responses {
			response := &models.ChannelResponse{
				CheckerNumber: data.CheckerNumber,
				Gr:            data.Gr / refPatchGrGb,
				Gb:            data.Gb / refPatchGrGb,
				R:             data.R / refPatchGrGb,
				B:             data.B / refPatchGrGb,
			}
			// stacking
			responses = append(responses, *response)
		}
	}

	return status, responses
}

// calculate
func (rc *responseController) calculateEachChannelResponse(ill map[int]float64, gr map[int]float64, gb map[int]float64, r map[int]float64, b map[int]float64, checker map[int]float64, wavelength int) (grch, gbch, rch, bch float64) {

	// calculate response
	grChRes := ill[wavelength] * gr[wavelength] * checker[wavelength]
	gbChRes := ill[wavelength] * gb[wavelength] * checker[wavelength]
	rChRes := ill[wavelength] * r[wavelength] * checker[wavelength]
	bChRes := ill[wavelength] * b[wavelength] * checker[wavelength]

	return grChRes, gbChRes, rChRes, bChRes
}