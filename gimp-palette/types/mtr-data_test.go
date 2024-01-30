package types

import (
	"testing"

	"gotest.tools/v3/assert"
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

	// check category name
	assert.Equal(t, category.Name, MTRLogoCategoryName)

	// check colors
	assert.DeepEqual(t, category.Colors, []*Color{
		// must be sorted by name
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
	})
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
