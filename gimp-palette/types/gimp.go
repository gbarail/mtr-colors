package types

import "fmt"

type GIMPPaletteCategory struct {
	Name   string
	Colors []*Color
}

type GIMPPalette struct {
	Name       string
	Categories []*GIMPPaletteCategory
}

func (c GIMPPaletteCategory) String() string {
	bytes := []byte{}

	// Write category name at the top, if it exists
	if c.Name != "" {
		bytes = fmt.Appendf(bytes, "# %s", c.Name)
		if len(c.Colors) > 0 {
			bytes = fmt.Append(bytes, "\n")
		}
	}

	// Write colors, one on each line
	for i, color := range c.Colors {
		rgb := color.RGB

		bytes = fmt.Appendf(bytes, "%3d %3d %3d", rgb[0], rgb[1], rgb[2])
		if color.Name != "" {
			bytes = fmt.Append(bytes, " ", color.Name)
		}
		if i < len(c.Colors)-1 {
			bytes = fmt.Append(bytes, "\n")
		}
	}

	return string(bytes)
}

func (p GIMPPalette) String() string {
	bytes := []byte{}

	// Write top
	bytes = fmt.Appendf(bytes, "GIMP Palette\nName: %s", p.Name)

	// Write each category
	for _, category := range p.Categories {
		bytes = fmt.Append(bytes, "\n\n", category)
	}

	return string(bytes)
}
