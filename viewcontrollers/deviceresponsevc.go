package viewcontrollers

import (
	"PixelTool/controllers"
	"PixelTool/models"
	"PixelTool/util"
)

/*
DeviceResponseViewController :device response view controller
*/
type DeviceResponseViewController interface {
	ReadResponseRawData(filepath map[string]string) bool

	CalculateDeviceResponse(ill models.IlluminationCode, start, stop, step int, gammma float64, refPatchNum int) bool
	CalculateLinearMatrix(elm []float64) bool
	CalculateWhiteBalanceGain(refPatchNumber int) (redGain, blueGain float64)
	Calculate8bitResponse(patchNumber int, data []float64, redGain, blueGain float64, refLevel uint8) *models.ColorCode

	// getters
	RawData() []models.ChannelResponse
	RawResponseData() []models.ChannelResponse
	LinearizedResponseData() [][]float64

	// steam out PNG patch image
	CreateColorCodePatch(data *models.ColorCode, filesavepath, dirname string, width, height int) bool
}

// defintion of structure
type deviceResponseViewController struct {
	resCon controllers.ResponseController // response controller

	// stockers
	rawData         []models.ChannelResponse
	rawResponseData []models.ChannelResponse

	linearizedResData [][]float64
}

/*
NewDeviceResponseViewController : initializer of VC
*/
func NewDeviceResponseViewController() DeviceResponseViewController {
	obj := new(deviceResponseViewController)

	// initialize properties
	obj.resCon = controllers.NewResponseController()

	// initialize stockers
	obj.rawData = make([]models.ChannelResponse, 0)
	obj.rawResponseData = make([]models.ChannelResponse, 0)
	obj.linearizedResData = make([][]float64, 0)

	return obj
}

/*
ReadResponseRawData	:
	in	;filepath map[string]string
	out	;bool
*/
func (vc *deviceResponseViewController) ReadResponseRawData(filepath map[string]string) bool {
	status := false

	// read response raw data
	if len(filepath) != 0 {
		status = vc.resCon.ReadResponseData(filepath)
	}

	return status
}

/*
CalculateDeviceResponse
	in	;
		ill models.Illumination,
		start 	;scan start wavelength
		stop	;scan stop wavelength
		step	;scan step wavelength
		refPatchNum
		refPatchLevel
	out	;bool
*/
func (vc *deviceResponseViewController) CalculateDeviceResponse(ill models.IlluminationCode, start, stop, step int, gamma float64, refPatchNum int) bool {
	status := false

	/*
		start to calculate channel response
			1. calculate channel response
			2. calculate gamma correction
			3. check slice size, and then update result to stocker
	*/
	process, responses := vc.resCon.CalculateChannelResponse(ill, start, stop, step, refPatchNum)
	if process {

		// stock rawdata
		vc.rawData = responses

		// initialize gamma correction buffer
		gammaCorrectedRes := make([]models.ChannelResponse, 0)

		// calculate gamma correction
		for _, data := range responses {
			flag, result := vc.resCon.CalculateGammaCorrection(gamma, &data)
			if flag {
				gammaCorrectedRes = append(gammaCorrectedRes, *result)
			}
		}

		// check slice size
		if len(gammaCorrectedRes) != 0 {
			vc.rawResponseData = gammaCorrectedRes

			//update status
			status = true
		}
	}

	return status
}

/*
CalculateLinearMatrix
	in	;elm []float64, grgbrb []float64
	out	;bool
*/
func (vc *deviceResponseViewController) CalculateLinearMatrix(elm []float64) bool {
	status := false

	if len(vc.rawResponseData) != 0 {
		// buffer
		responses := make([][]float64, 0)
		for _, data := range vc.rawResponseData {
			// change data format to linear matrix calculation
			grgbrb := []float64{
				data.Gr,
				data.Gb,
				data.R,
				data.B,
			}

			// calcualte linear matrix
			response := vc.resCon.CalculateLinearMatrix(elm, grgbrb)

			// stock
			responses = append(responses, response)
		}

		// check response
		if len(responses) != 0 {
			vc.linearizedResData = responses

			// update status
			status = true
		}
	}

	return status
}

/*
CalculateWhiteBalanceGain
	in	;refPtchNumber int
	out	;redGain, blueGain float64
*/
func (vc *deviceResponseViewController) CalculateWhiteBalanceGain(refPatchNumber int) (redGain, blueGain float64) {
	if refPatchNumber > -1 && refPatchNumber < 25 {
		return vc.resCon.CalculateWhiteBalanceGain(vc.linearizedResData[refPatchNumber-1])
	} else {
		return 0.0, 0.0
	}
}

/*
Calculate8bitResponse
	in	;data []float64, redGain, blueGain float64, refLevel uint8
	out	;models.ColorCode
*/
func (vc *deviceResponseViewController) Calculate8bitResponse(patchNumber int, data []float64, redGain, blueGain float64, refLevel uint8) *models.ColorCode {

	// calculate raw data
	red := data[0] * redGain
	green := data[1]
	blue := data[2] * blueGain

	// digitize signal
	digitizer := util.NewDigitizer()
	red8bit := digitizer.D8bitDigitizeData(red, refLevel)
	green8bit := digitizer.D8bitDigitizeData(green, refLevel)
	blue8bit := digitizer.D8bitDigitizeData(blue, refLevel)

	// patch name
	pname := models.MacbethColorCode(patchNumber).String()

	// create color code model
	colorcode := models.SetColorCode(patchNumber+1, pname, red8bit, green8bit, blue8bit, 255)

	return colorcode

}

/*
CreateColorCodePatch
	in	;filesavepath, dirname string, width, height int
	out	;bool
*/
func (vc *deviceResponseViewController) CreateColorCodePatch(data *models.ColorCode, filesavepath, dirname string, width, height int) bool {
	status := false

	if filesavepath != "" && dirname != "" && data != nil {
		/*
			TODO	;impliment stream out
		*/
	}

	return status
}

/*
RawData
	out	;[]models.ChannelResponse
*/
func (vc *deviceResponseViewController) RawData() []models.ChannelResponse {
	return vc.rawData
}

/*
RawResponseData
	out	;[]models.ChannelResponse
*/
func (vc *deviceResponseViewController) RawResponseData() []models.ChannelResponse {
	return vc.rawResponseData
}

/*
LinearizedResponseData
	out ;[][]float64
*/
func (vc *deviceResponseViewController) LinearizedResponseData() [][]float64 {
	return vc.linearizedResData
}
