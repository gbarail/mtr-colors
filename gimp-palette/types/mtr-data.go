package types

type MTRLogoData map[string]*Color

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
