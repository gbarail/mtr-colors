package main

import (
	"fmt"

	"github.com/gbarail/mtr-colors/types"
)

func main() {
	palette := types.GIMPPalette{
		Name: "Hong Kong MTR colors",
	}

	{ // Generate MTR Logo category
		data, err := ReadMTRLogoYAMLFile()
		if err != nil {
			panic(err)
		}

		category := types.GenerateMTRLogoCategory(data)
		palette.Categories = append(palette.Categories, category)
	}

	{ // Generate MTR System Map categories
		data, err := ReadMTRSystemMapYAMLFile()
		if err != nil {
			panic(err)
		}

		categories := types.GenerateMTRSystemMapCategories(data)
		palette.Categories = append(palette.Categories, categories...)
	}

	fmt.Println(palette)
}
