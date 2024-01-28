package types

import (
	"fmt"
	"strings"
)

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
		bytes = fmt.Appendf(bytes, "# %s\n", c.Name)
	}

	// Write colors, one on each line
	for i, color := range c.Colors {
		rgb := color.RGB

		bytes = fmt.Appendf(bytes, "%3d %3d %3d", rgb[0], rgb[1], rgb[2])
		if color.Name != "" {
			bytes = fmt.Appendf(bytes, " # %s", color.Name)
		}
		if i < len(c.Colors)-1 {
			bytes = fmt.Appendf(bytes, "\n")
		}
	}

	return string(bytes)
}

func (p GIMPPalette) String() string {
	data := []string{}

	// Output top
	top := fmt.Sprintf("GIMP Palette\nName: %s", p.Name)
	data = append(data, top)

	for _, category := range p.Categories {
		data = append(data, "\n\n")
		data = append(data, category.String())
	}

	return strings.Join(data, "")
}
