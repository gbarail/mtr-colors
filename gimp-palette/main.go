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
		palette.Categories = append(palette.Categories, &category)
	}

	fmt.Printf("%v\n", palette)
}
