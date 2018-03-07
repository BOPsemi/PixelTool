package models

import "testing"
import "github.com/stretchr/testify/assert"

func Test_New(t *testing.T) {
	obj := new(ColorCode)

	assert.NotNil(t, obj)
}

func Test_ReadColorCode(t *testing.T) {
	path := "/Users/kazufumiwatanabe/go/src/PixelTool/data/macbeth_patch_code.csv"
	colorCodes := ReadColorCode(path)

	assert.EqualValues(t, 24, len(colorCodes))
}
