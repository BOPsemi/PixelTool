package controllers

import "image/color"
import "image"

/*
Image controller package
*/

/*
ImageController :interface of image controller
*/
type ImageController interface {
	// create raw image data
	CreateImage(data []color.RGBA, height, width int) *image.RGBA
	CreateSolidImage(data color.RGBA, height, width int) *image.RGBA

	// read pixel data from image file
	SerializeImage(img *image.RGBA) []color.RGBA
}

// definition of image controller
type imageContrller struct {
}

// NewImageController :initializer of image controller
func NewImageController() ImageController {
	obj := new(imageContrller)

	return obj
}

/*
CreateImage :create image from data
	data 				<-[]color.RGBA
	height, width 		<-int
*/
func (im *imageContrller) CreateImage(data []color.RGBA, height, width int) *image.RGBA {
	img := new(image.RGBA)

	if height > 0 && width > 0 {
		// check data size
		if (height * width) == len(data) {

			// create image
			canvas := image.NewRGBA(image.Rect(0, 0, width, height))
			for i := 0; i < width; i++ {
				for j := 0; j < height; j++ {
					index := width*i + j

					// raw data
					rawData := color.RGBA{
						R: data[index].R,
						G: data[index].G,
						B: data[index].B,
						A: 255,
					}

					// draw the raw data on canvas
					canvas.Set(i, j, rawData)
				}
			}

			// update image
			img = canvas
		}
	}
	return img
}

/*
CreateSolidImage(data color.RGBA, height, width int) *image.RGBA
*/
func (im *imageContrller) CreateSolidImage(data color.RGBA, height, width int) *image.RGBA {
	img := new(image.RGBA)

	if height > 0 && width > 0 {

		// create image
		canvas := image.NewRGBA(image.Rect(0, 0, width, height))
		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {

				// draw the raw data on canvas
				canvas.Set(i, j, data)
			}
		}

		// update image
		img = canvas

	}

	return img
}

/*
SerializeImage :serialize image data to color.RGBA slice
	img 	:*image.RGBA
*/
func (im *imageContrller) SerializeImage(img *image.RGBA) []color.RGBA {
	data := make([]color.RGBA, 0)

	if img != nil {
		for i := 0; i < img.Bounds().Size().X; i++ {
			for j := 0; j < img.Bounds().Size().Y; j++ {

				// extract point data
				rgba := img.At(i, j)

				// each channel data
				r, g, b, a := rgba.RGBA()

				// create raw data
				rawdata := color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(a),
				}

				// stack data
				data = append(data, rawdata)
			}
		}
	}

	return data
}
