package models

/*
ColorSpace :enum for light source
*/
type ColorSpace int

const (
	CIE ColorSpace = iota
	NTSC
	SRGB
)
