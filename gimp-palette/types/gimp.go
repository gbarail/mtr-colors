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
	data := []string{}

	// Output top
	top := fmt.Sprintf("GIMP Palette\nName: %s", p.Name)
	data = append(data, top)

	for _, category := range p.Categories {
		data = append(data, "\n\n")

		// Output category name if it exists
		if category.Name != "" {
			data = append(data, fmt.Sprintf("# %s\n", category.Name))
		}

		// Output colors
		lines := []string{}
		for _, color := range category.Colors {
			RGB := color.RGB

			end := ""
			if color.Name != "" {
				end = fmt.Sprintf(" # %s", color.Name)
			}

			line := fmt.Sprintf("%3d %3d %3d%s", RGB[0], RGB[1], RGB[2], end)
			lines = append(lines, line)
		}
		data = append(data, strings.Join(lines, "\n"))
	}

	return strings.Join(data, "")
}
