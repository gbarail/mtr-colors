package types

type Category struct {
	Name   string
	Colors []*Color
}

type GIMPPalette []*Category
