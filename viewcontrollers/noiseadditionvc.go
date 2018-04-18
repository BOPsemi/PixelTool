package viewcontrollers

import (
	"PixelTool/controllers"
	"PixelTool/util"
	"image"
	"image/color"
)

/*
NoiseAdditionViewController :interface of noise addtion vc
*/
type NoiseAdditionViewController interface {
	SetImageDataForWhitePixelAddition(basefilepath, noisefilepath string) bool
}

// definition of structure
type noiseAdditionViewController struct {
	base  []color.RGBA // base image data
	noise []color.RGBA // noise image data
}

/*
NewNoiseAdditionViewController :initializer of noise addtion vc
*/
func NewNoiseAdditionViewController() NoiseAdditionViewController {
	obj := new(noiseAdditionViewController)

	// init properties
	obj.base = make([]color.RGBA, 0)
	obj.noise = make([]color.RGBA, 0)

	return obj
}

/*
SetImageDataForWhitePixelAddition :
	in	;basefilepath, noisefilepath string
	out	;bool
*/
func (vc *noiseAdditionViewController) SetImageDataForWhitePixelAddition(basefilepath, noisefilepath string) bool {
	status := false

	if basefilepath != "" && noisefilepath != "" {

		// read base image data
		baseData := vc.imageFileOpen(basefilepath)
		if len(baseData) > 0 {
			vc.base = baseData
		}

		// read noise image data
		noiseData := vc.imageFileOpen(noisefilepath)
		if len(noiseData) > 0 {
			vc.noise = noiseData
		}
	}

	return status
}

func (vc *noiseAdditionViewController) imageFileOpen(filepath string) []color.RGBA {
	data := make([]color.RGBA, 0)

	if filepath != "" {
		iohandler := util.NewIOUtil()
		imageData := iohandler.ReadImageFile(filepath)

		if imageData != nil {
			// type change
			if img, ok := imageData.(*image.RGBA); ok {

				// call iamge controller
				imagecontroller := controllers.NewImageController()
				rgbaImagaData := imagecontroller.SerializeImage(img)

				// check data size
				if len(rgbaImagaData) > 0 {
					// update data
					data = rgbaImagaData
				}
			}
		}
	}

	return data
}
