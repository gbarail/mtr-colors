package types

type RGBValue uint8
type RGB [3]RGBValue

type Color struct {
	Name string
	RGB  RGB
}

func ColorCompare(a, b *Color) int {
	switch {
	case a.Name < b.Name:
		return -1
	case a.Name > b.Name:
		return 1
	default:
		return 0
	}
}
