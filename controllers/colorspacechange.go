package controllers

import (
	"PixelTool/models"

	"gonum.org/v1/gonum/mat"
)

/*
ColorSpaceChange :interface of color space change
*/
type ColorSpaceChange interface {
	ReadWhitePoint(path string) bool

	SpaceChangeRGBtoXYZ(colorSP models.ColorSpace) []float64
	SpaceChangeXYZtoRGB(colorSP models.ColorSpace) []float64
}

// definition of colorSpaceChange
type colorSpaceChange struct {
	whitePoint *models.WhitePoint
}

/*
NewColorSpaceChange :initializer of color space change
*/
func NewColorSpaceChange() ColorSpaceChange {
	obj := new(colorSpaceChange)

	return obj
}

/*
ReadWhitePoint
	in	;path string	// JSON file path
	out	;bool			// status updte
*/
func (cs *colorSpaceChange) ReadWhitePoint(path string) bool {
	status := false
	if whitepoint, ok := models.ReadWhitePoint(path); ok {
		// update whitePoint data
		cs.whitePoint = whitepoint

		// update status
		status = true
	}

	return status
}

/*
SpaceChangeRGBtoXYZ :
	in	;colorSP models.ColorSpace
	out	;[]float64
*/
func (cs *colorSpaceChange) SpaceChangeRGBtoXYZ(colorSP models.ColorSpace) []float64 {
	result := make([]float64, 0)

	// init white point
	rgbElm, wElm := cs.initWhitePointMatrix(colorSP)

	if len(rgbElm) == 9 && len(wElm) == 3 {
		// create matrix from data
		rgbElmMat := mat.NewDense(3, 3, rgbElm)
		wElmMat := mat.NewDense(3, 1, wElm)

		// make skelton
		invRgbElmMat := mat.NewDense(3, 3, make([]float64, 9)) // for inversed mat
		wmat2 := mat.NewDense(3, 1, make([]float64, 3))        // for internal calc
		matResult := mat.NewDense(3, 3, make([]float64, 9))    // result

		// calculate inversed matrix of rgbElmMat
		invRgbElmMat.Inverse(rgbElmMat)
		invRgbElmMat.Inverse(rgbElmMat)  // calculate inversed matrix
		wmat2.Mul(invRgbElmMat, wElmMat) // M-1 x w

		// create transversal rgb matrix
		trgb := []float64{
			wmat2.At(0, 0), 0.0, 0.0,
			0.0, wmat2.At(1, 0), 0.0,
			0.0, 0.0, wmat2.At(2, 0),
		}

		trgbMat := mat.NewDense(3, 3, trgb)
		/*
			|	fx	0	0 	|
			|	0	fy	0	|
			|	0	0	fz	|
		*/
		// finalize calculation
		matResult.Mul(rgbElmMat, trgbMat) // rgbElm x trgb
		result = []float64{
			matResult.At(0, 0), matResult.At(0, 1), matResult.At(0, 2),
			matResult.At(1, 0), matResult.At(1, 1), matResult.At(1, 2),
			matResult.At(2, 0), matResult.At(2, 1), matResult.At(2, 2),
		}
	}

	return result
}

/*
SpaceChangeXYZtoRGB	:
	in	;colorSP models.ColorSpace
	out	;[]float64
*/
func (cs *colorSpaceChange) SpaceChangeXYZtoRGB(colorSP models.ColorSpace) []float64 {
	result := make([]float64, 0)

	return result
}

// initialize white point matrix
func (cs *colorSpaceChange) initWhitePointMatrix(colorSP models.ColorSpace) (rgb []float64, w []float64) {
	switch colorSP {
	case models.CIE:
		/*
			CIE Color Space Case
		*/
		// calculate Z elements
		cs.whitePoint.CIE.R.Z = 1.0 - cs.whitePoint.CIE.R.X - cs.whitePoint.CIE.R.Y
		cs.whitePoint.CIE.G.Z = 1.0 - cs.whitePoint.CIE.G.X - cs.whitePoint.CIE.G.Y
		cs.whitePoint.CIE.B.Z = 1.0 - cs.whitePoint.CIE.B.X - cs.whitePoint.CIE.B.Y

		// calculate white point matxrix
		cs.whitePoint.CIE.W.Z2 = 1.0 - cs.whitePoint.CIE.W.X2 - cs.whitePoint.CIE.W.Y2

		// calculate XYZ elements
		cs.whitePoint.CIE.Wn.X = cs.whitePoint.CIE.W.X2 / cs.whitePoint.CIE.W.Y2
		cs.whitePoint.CIE.Wn.Y = cs.whitePoint.CIE.W.Y2 / cs.whitePoint.CIE.W.Y2
		cs.whitePoint.CIE.Wn.Z = cs.whitePoint.CIE.W.Z2 / cs.whitePoint.CIE.W.Y2

		// make slice for RGB
		rgbElm := []float64{
			cs.whitePoint.CIE.R.X, cs.whitePoint.CIE.G.X, cs.whitePoint.CIE.B.X,
			cs.whitePoint.CIE.R.Y, cs.whitePoint.CIE.G.Y, cs.whitePoint.CIE.B.Y,
			cs.whitePoint.CIE.R.Z, cs.whitePoint.CIE.G.Z, cs.whitePoint.CIE.B.Z,
		}

		// make slice for white point
		wElm := []float64{
			cs.whitePoint.CIE.Wn.X, cs.whitePoint.CIE.Wn.Y, cs.whitePoint.CIE.Wn.Z,
		}

		// return the calculated results
		return rgbElm, wElm

	case models.SRGB:
		/*
			sRGB Color Space Case
		*/
		// calculate Z elements
		cs.whitePoint.SRGB.R.Z = 1.0 - cs.whitePoint.SRGB.R.X - cs.whitePoint.SRGB.R.Y
		cs.whitePoint.SRGB.G.Z = 1.0 - cs.whitePoint.SRGB.G.X - cs.whitePoint.SRGB.G.Y
		cs.whitePoint.SRGB.B.Z = 1.0 - cs.whitePoint.SRGB.B.X - cs.whitePoint.SRGB.B.Y

		// calculate white point matxrix
		cs.whitePoint.SRGB.W.Z2 = 1.0 - cs.whitePoint.SRGB.W.X2 - cs.whitePoint.SRGB.W.Y2

		// calculate XYZ elements
		cs.whitePoint.SRGB.Wn.X = cs.whitePoint.SRGB.W.X2 / cs.whitePoint.SRGB.W.Y2
		cs.whitePoint.SRGB.Wn.Y = cs.whitePoint.SRGB.W.Y2 / cs.whitePoint.SRGB.W.Y2
		cs.whitePoint.SRGB.Wn.Z = cs.whitePoint.SRGB.W.Z2 / cs.whitePoint.SRGB.W.Y2

		// make slice for RGB
		rgbElm := []float64{
			cs.whitePoint.SRGB.R.X, cs.whitePoint.SRGB.G.X, cs.whitePoint.SRGB.B.X,
			cs.whitePoint.SRGB.R.Y, cs.whitePoint.SRGB.G.Y, cs.whitePoint.SRGB.B.Y,
			cs.whitePoint.SRGB.R.Z, cs.whitePoint.SRGB.G.Z, cs.whitePoint.SRGB.B.Z,
		}

		// make slice for white point
		wElm := []float64{
			cs.whitePoint.SRGB.Wn.X, cs.whitePoint.SRGB.Wn.Y, cs.whitePoint.SRGB.Wn.Z,
		}

		// return the calculated results
		return rgbElm, wElm

	case models.NTSC:
		/*
			NTSC Color Space Case
		*/
		// calculate Z elements
		cs.whitePoint.NTSC.R.Z = 1.0 - cs.whitePoint.NTSC.R.X - cs.whitePoint.NTSC.R.Y
		cs.whitePoint.NTSC.G.Z = 1.0 - cs.whitePoint.NTSC.G.X - cs.whitePoint.NTSC.G.Y
		cs.whitePoint.NTSC.B.Z = 1.0 - cs.whitePoint.NTSC.B.X - cs.whitePoint.NTSC.B.Y

		// calculate white point matxrix
		cs.whitePoint.NTSC.W.Z2 = 1.0 - cs.whitePoint.NTSC.W.X2 - cs.whitePoint.NTSC.W.Y2

		// calculate XYZ elements
		cs.whitePoint.NTSC.Wn.X = cs.whitePoint.NTSC.W.X2 / cs.whitePoint.NTSC.W.Y2
		cs.whitePoint.NTSC.Wn.Y = cs.whitePoint.NTSC.W.Y2 / cs.whitePoint.NTSC.W.Y2
		cs.whitePoint.NTSC.Wn.Z = cs.whitePoint.NTSC.W.Z2 / cs.whitePoint.NTSC.W.Y2

		// make slice for RGB
		rgbElm := []float64{
			cs.whitePoint.NTSC.R.X, cs.whitePoint.NTSC.G.X, cs.whitePoint.NTSC.B.X,
			cs.whitePoint.NTSC.R.Y, cs.whitePoint.NTSC.G.Y, cs.whitePoint.NTSC.B.Y,
			cs.whitePoint.NTSC.R.Z, cs.whitePoint.NTSC.G.Z, cs.whitePoint.NTSC.B.Z,
		}

		// make slice for white point
		wElm := []float64{
			cs.whitePoint.NTSC.Wn.X, cs.whitePoint.NTSC.Wn.Y, cs.whitePoint.NTSC.Wn.Z,
		}

		// return the calculated results
		return rgbElm, wElm

	default:
		return []float64{}, []float64{}
	}
}
