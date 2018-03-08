package models

import (
	"fmt"
	"testing"
)

func Test_ReadColorChecker(t *testing.T) {
	path := "/Users/kazufumiwatanabe/go/src/PixelTool/data/Macbeth_Color_Checker.csv"
	colorcheckers := ReadColorChecker(path)

	fmt.Println(colorcheckers[1])
}
