package types

type RGBValue uint8
type RGB [3]RGBValue

type Color struct {
	Name string
	RGB  RGB
}
