package types

type RGBValue uint8
type RGB [3]RGBValue

type Color struct {
	Name string
	RGB  RGB
}

func CompareColors(a, b *Color) int {
	if c := compareColorNames(a, b); c != 0 {
		return c
	}
	return compareRGB(&a.RGB, &b.RGB)
}

func compareColorNames(a, b *Color) int {
	switch {
	case a.Name < b.Name:
		return -1
	case a.Name > b.Name:
		return 1
	default:
		return 0
	}
}

func compareRGB(a, b *RGB) int {
	for i, va := range a {
		vb := b[i]
		if c := compareRGBValues(va, vb); c != 0 {
			return c
		}
	}
	return 0
}

func compareRGBValues(a, b RGBValue) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
