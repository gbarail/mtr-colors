package main

import (
	"fmt"

	"github.com/gbarail/mtr-colors/types"
)

func main() {
	palette := types.GIMPPalette{
		Name: "Hong Kong MTR colors",
	}

	data, err := ReadMTRLogoYAMLFile()
	if err != nil {
		panic(err)
	}

	category := types.Category{
		Name:   "MTR logo",
		Colors: []*types.Color{},
	}

	for k, v := range data {
		category.Colors = append(category.Colors, &types.Color{
			Name: k,
			RGB:  v.RGB,
		})
	}

	palette.Categories = append(palette.Categories, &category)

	fmt.Printf("%v\n", palette)
}
