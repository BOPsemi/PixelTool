package controllers

import (
	"PixelTool/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	PATH = "/Users/kazufumiwatanabe/go/src/PixelTool/json/whitepoint.json"
)

func Test_NewColorSpaceChange(t *testing.T) {
	obj := NewColorSpaceChange()

	assert.NotNil(t, obj)
}

func Test_ReadWhitePoint(t *testing.T) {

	obj := NewColorSpaceChange()
	assert.True(t, obj.ReadWhitePoint(PATH))

}

func Test_SpaceChangeRGBtoXYZ(t *testing.T) {
	obj := NewColorSpaceChange()
	if obj.ReadWhitePoint(PATH) {
		rgb := []float64{1.0, 0.0, 0.0}

		// SRGB
		xyz := obj.SpaceChangeRGBtoXYZ(models.SRGB, rgb)
		lab := obj.SpaceChangeXYZtoLab(models.SRGB, xyz)
		fmt.Println(lab)

		// CIW
		xyz = obj.SpaceChangeRGBtoXYZ(models.CIE, rgb)
		lab = obj.SpaceChangeXYZtoLab(models.CIE, xyz)
		fmt.Println(lab)

		// NTSC
		xyz = obj.SpaceChangeRGBtoXYZ(models.NTSC, rgb)
		lab = obj.SpaceChangeXYZtoLab(models.NTSC, xyz)
		fmt.Println(lab)
	}

}
