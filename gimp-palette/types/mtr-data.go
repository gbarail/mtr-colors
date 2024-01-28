package types

type MTRLogoData map[string]*Color

type MTRSystemMapData struct {
	LinesColors map[string]*struct {
		FullName string `yaml:"full_name"`
		Color    *Color
	} `yaml:"lines_colors"`
	MiscellaneousColors map[string]*struct {
		Name  string
		Color *Color
	} `yaml:"miscellaneous_colors"`
}

func GenerateMTRLogoCategory(data MTRLogoData) GIMPPaletteCategory {
	category := GIMPPaletteCategory{
		Name: "MTR logo",
	}

	for _, v := range data {
		category.Colors = append(category.Colors, v)
	}

	return category
}

func GenerateMTRSystemMapCategories(data *MTRSystemMapData) []*GIMPPaletteCategory {
	var categories []*GIMPPaletteCategory

	{ // lines colors
		category := &GIMPPaletteCategory{
			Name: "Lines colors",
		}

		for _, v := range data.LinesColors {
			category.Colors = append(category.Colors, &Color{
				Name: v.FullName, // full name of the line
				RGB:  v.Color.RGB,
			})
		}

		categories = append(categories, category)
	}

	{ // miscellaneous colors
		category := &GIMPPaletteCategory{
			Name: "Miscellaneous colors",
		}

		for _, v := range data.MiscellaneousColors {
			category.Colors = append(category.Colors, &Color{
				Name: v.Name,
				RGB:  v.Color.RGB,
			})
		}

		categories = append(categories, category)
	}

	return categories
}
