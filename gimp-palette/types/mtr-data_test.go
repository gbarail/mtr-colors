package types

import (
	"testing"
)

func TestGenerateMTRLogoCategory(t *testing.T) {
	testData := MTRLogoData{
		"red": {
			Name: "red",
			RGB:  RGB{255, 0, 0},
		},
		"green": {
			Name: "green",
			RGB:  RGB{0, 255, 0},
		},
		"blue": {
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

func TestGenerateMTRSystemMapCategories(t *testing.T) {
	testData := &MTRSystemMapData{
		LinesColors: map[string]*LineColor{
			"red": {
				FullName: "Red",
				Color: &Color{
					RGB: RGB{255, 0, 0},
				},
			},
			"green": {
				FullName: "Green",
				Color: &Color{
					RGB: RGB{0, 255, 0},
				},
			},
			"blue": {
				FullName: "Blue",
				Color: &Color{
					RGB: RGB{0, 0, 255},
				},
			},
		},
		MiscellaneousColors: map[string]*MiscellaneousColor{
			"background": {
				Name: "Background",
				Color: &Color{
					RGB: RGB{255, 255, 255},
				},
			},
			"text": {
				Name: "Text",
				Color: &Color{
					RGB: RGB{0, 0, 0},
				},
			},
		},
	}

	categories := GenerateMTRSystemMapCategories(testData)

	expectedCategories := []*GIMPPaletteCategory{
		{
			Name: "Lines colors",
			Colors: []*Color{ // sorted by names
				{
					Name: "Blue",
					RGB:  RGB{0, 0, 255},
				},
				{
					Name: "Green",
					RGB:  RGB{0, 255, 0},
				},
				{
					Name: "Red",
					RGB:  RGB{255, 0, 0},
				},
			},
		},
		{
			Name: "Miscellaneous colors",
			Colors: []*Color{ // sorted by names
				{
					Name: "Background",
					RGB:  RGB{255, 255, 255},
				},
				{
					Name: "Text",
					RGB:  RGB{0, 0, 0},
				},
			},
		},
	}

	checkCategories := func(actual, expected *GIMPPaletteCategory) {
		if actual.Name != expected.Name {
			t.Errorf("Expected category name to be \"%s\", got \"%s\"",
				expected.Name, actual.Name)
		}
		if len(actual.Colors) != len(expected.Colors) {
			t.Errorf("Expected %d colors, got %d colors",
				len(expected.Colors), len(actual.Colors))
		}

		for i, color := range actual.Colors {
			if *color != *expected.Colors[i] {
				t.Errorf("Expected color %d to be %v, got %v",
					i, expected.Colors[i], color)
			}
		}
	}

	{ // check number of categories
		if len(categories) != len(expectedCategories) {
			t.Errorf("Expected %d categories, got %d categories",
				len(expectedCategories), len(categories))
		}
	} // check each category
	{
		for i, category := range categories {
			checkCategories(category, expectedCategories[i])
		}
	}
}
