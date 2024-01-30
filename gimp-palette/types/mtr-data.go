package types

import "slices"

type MTRLogoData map[string]*Color

type MTRSystemMapData struct {
	LinesColors         map[string]*LineColor `yaml:"lines_colors"`
	MiscellaneousColors map[string]*MiscellaneousColor
}

type LineColor struct {
	FullName string `yaml:"full_name"`
	Color    *Color
}

type MiscellaneousColor struct {
	Name  string
	Color *Color
}

const (
	MTRLogoCategoryName = "MTR logo"

	LinesColorsCategoryName         = "Lines colors"
	MiscellaneousColorsCategoryName = "Miscellaneous colors"
)

func GenerateMTRLogoCategory(data MTRLogoData) *GIMPPaletteCategory {
	category := &GIMPPaletteCategory{
		Name: MTRLogoCategoryName,
	}

	for _, v := range data {
		category.Colors = append(category.Colors, v)
	}
	slices.SortFunc(category.Colors, CompareColors)

	return category
}

func GenerateMTRSystemMapCategories(data *MTRSystemMapData) []*GIMPPaletteCategory {
	var categories []*GIMPPaletteCategory

	{ // lines colors
		category := &GIMPPaletteCategory{
			Name: LinesColorsCategoryName,
		}

		for _, v := range data.LinesColors {
			category.Colors = append(category.Colors, &Color{
				Name: v.FullName, // full name of the line
				RGB:  v.Color.RGB,
			})
		}
		slices.SortFunc(category.Colors, CompareColors)

		categories = append(categories, category)
	}

	{ // miscellaneous colors
		category := &GIMPPaletteCategory{
			Name: MiscellaneousColorsCategoryName,
		}

		for _, v := range data.MiscellaneousColors {
			category.Colors = append(category.Colors, &Color{
				Name: v.Name,
				RGB:  v.Color.RGB,
			})
		}
		slices.SortFunc(category.Colors, CompareColors)

		categories = append(categories, category)
	}

	return categories
}
