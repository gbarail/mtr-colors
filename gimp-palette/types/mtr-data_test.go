package types

import (
	"testing"
)

func TestGenerateMTRLogoCategory(t *testing.T) {
	testData := MTRLogoData{
		"red": &Color{
			Name: "red",
			RGB:  RGB{255, 0, 0},
		},
		"green": &Color{
			Name: "green",
			RGB:  RGB{0, 255, 0},
		},
		"blue": &Color{
			Name: "blue",
			RGB:  RGB{0, 0, 255},
		},
	}

	category := GenerateMTRLogoCategory(testData)
	{ // check category name
		expectedCategoryName := "MTR logo"
		if category.Name != expectedCategoryName {
			t.Errorf("Expected category name to be \"%s\", got \"%s\"",
				expectedCategoryName, category.Name)
		}
	}
	{ // check colors
		expectedColors := []Color{
			// must be sorted
			{
				Name: "blue",
				RGB:  RGB{0, 0, 255},
			},
			{
				Name: "green",
				RGB:  RGB{0, 255, 0},
			},
			{
				Name: "red",
				RGB:  RGB{255, 0, 0},
			},
		}

		for i, color := range category.Colors {
			if *color != expectedColors[i] {
				t.Errorf("Expected color %d to be %v, got %v",
					i, expectedColors[i], color)
			}
		}
	}
}
