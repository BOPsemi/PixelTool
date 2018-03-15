package controllers

import (
	"encoding/json"
	"io/ioutil"
)

/*
StartUpController :interface of Start up controller structure
*/
type StartUpController interface {
	InitRawdata(path string) (Rawdata, bool)
}

// Rawdata :definition of raw data str
type Rawdata struct {
	DeviceQE        string `json:"deviceQE"`
	IlluminationA   string `json:"illuminationA"`
	IlluminationD65 string `json:"illuminationD65"`
	ColorChecker    string `json:"colorChecker"`
	ColorCode       string `json:"colorCode"`
	WhitePixel      string `json:"whitePixel"`
}

// startUpController :structure definition
type startUpController struct {
	rawdata *Rawdata // raw data object
}

/*
NewStartUpController :initializer
*/
func NewStartUpController() StartUpController {
	obj := new(startUpController)

	// initialize objects
	obj.rawdata = new(Rawdata)

	return obj
}

/*
InitRawdata(path string) (Rawdata, bool)
*/

func (st *startUpController) InitRawdata(path string) (Rawdata, bool) {
	status := false

	if path != "" {
		if st.mapper(path) {
			status = true
		}
	} else {
		st.rawdata = nil
	}

	return *st.rawdata, status
}

// json mapper for RawData structure
func (st *startUpController) mapper(path string) bool {
	status := false

	// read json data
	data, err := ioutil.ReadFile(path)
	if err != nil {
		// failed to open read file
		st.rawdata = nil
	} else {
		// unmarshal
		err = json.Unmarshal(data, st.rawdata)
		if err != nil {
			// failed to mapping
			st.rawdata = nil
		} else {
			// successed to mapping
			status = true
		}
	}

	return status
}
