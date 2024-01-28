package types

type MTRLogoData map[string]*Color

type MTRSystemMapData struct {
	LinesColors map[string]*struct {
		FullName string `yaml:"full_name"`
		Color    *Color `yaml:"color"`
	} `yaml:"lines_colors"`
	MiscellaneousColors map[string]*struct {
		Name  string
		Color *Color
	} `yaml:"miscellaneous_colors"`
}

func GenerateMTRLogoCategory(data MTRLogoData) GIMPPaletteCategory {
	category := GIMPPaletteCategory{
		Name:   "MTR logo",
		Colors: []*Color{},
	}

	for _, v := range data {
		category.Colors = append(category.Colors, &Color{
			Name: v.Name,
			RGB:  v.RGB,
		})
	}

	return category
}

func GenerateMTRSystemMapCategories(data *MTRSystemMapData) []*GIMPPaletteCategory {
	categories := []*GIMPPaletteCategory{}

	{ // lines colors
		category := &GIMPPaletteCategory{
			Name:   "Lines colors",
			Colors: []*Color{},
		}

		for _, v := range data.LinesColors {
			category.Colors = append(category.Colors, v.Color)
		}

		categories = append(categories, category)
	}

	{ // miscellaneous colors
		category := &GIMPPaletteCategory{
			Name:   "Miscellaneous colors",
			Colors: []*Color{},
		}

		for _, v := range data.MiscellaneousColors {
			category.Colors = append(category.Colors, v.Color)
		}

		categories = append(categories, category)
	}

	return categories
}
