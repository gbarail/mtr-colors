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

func (p GIMPPalette) String() string {
	data := make([]string, 256)

	// Output top
	top := fmt.Sprintf("GIMP Palette\nName: %s\n", p.Name)
	data = append(data, top)

	for _, category := range p.Categories {
		data = append(data, "\n")

		// Output category name if it exists
		if category.Name != "" {
			data = append(data, fmt.Sprintf("# %s\n", category.Name))
		} else {
			data = append(data, "\n")
		}

		// Output colors
		for _, color := range category.Colors {
			RGB := color.RGB
			data = append(data, fmt.Sprintf("%3d %3d %3d", RGB[0], RGB[1], RGB[2]))

			if color.Name != "" {
				data = append(data, fmt.Sprintf(" # %s\n", color.Name))
			} else {
				data = append(data, "\n")
			}
		}
	}

	return strings.Join(data, "")
}
