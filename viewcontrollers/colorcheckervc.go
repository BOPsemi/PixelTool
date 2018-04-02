package viewcontrollers

import (
	"PixelTool/controllers"
	"PixelTool/models"
	"PixelTool/util"
)

/*
ColorCheckerViewController :control view
	- generate color checker patches
*/
type ColorCheckerViewController interface {
	CreateColorCodePatch(csvfilepath, filesavepath, dirname string, width, height int) bool
}

// strcture definition
type colorCheckerViewController struct {
	imgcontroller controllers.ImageController
	dirhandler    util.DirectoryHandler
	iohandler     util.IOUtil

	// properties
	colorCodes []models.ColorCode
}

/*
NewColorCheckerViewController : initializer
*/
func NewColorCheckerViewController() ColorCheckerViewController {
	obj := new(colorCheckerViewController)

	// initialize instances
	obj.imgcontroller = controllers.NewImageController()
	obj.dirhandler = util.NewDirectoryHandler()
	obj.iohandler = util.NewIOUtil()

	return obj
}

/*
CreateColorCodePatch(csvfilepath, filesavepath, dirname string) bool
*/
func (cc *colorCheckerViewController) CreateColorCodePatch(csvfilepath, filesavepath, dirname string, width, height int) bool {
	status := false

	if (csvfilepath != "") && (filesavepath != "") && (dirname != "") {
		// initialize data directory
		if cc.dirhandler.MakeDirectory(filesavepath, dirname) {
			// initalize colorcodes
			cc.colorCodes = models.ReadColorCode(csvfilepath)
			path := filesavepath + dirname + "/"

			// create solid images
			if len(cc.colorCodes) > 0 {
				for _, data := range cc.colorCodes {
					rawimage := cc.imgcontroller.CreateSolidImage(*data.GenerateColorRGBA(), width, height)
					cc.iohandler.StreamOutPNGFile(path, data.GetName(), rawimage)
				}

				// status update
				status = true
			}
		}
	}
	return status
}
