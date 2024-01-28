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

const GIMPPaletteTop = "GIMP Palette"

func (p GIMPPalette) String() string {
	data := make([]string, 256)

	top := fmt.Sprintf("%s\nName: %s\n\n", GIMPPaletteTop, p.Name)
	data = append(data, top)

	for _, category := range p.Categories {
		if category.Name != "" {
			data = append(data, fmt.Sprintf("# %s\n", category.Name))
		} else {
			data = append(data, "\n")
		}

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
