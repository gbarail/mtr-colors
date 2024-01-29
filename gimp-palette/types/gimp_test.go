package types

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, currFile, _, _ = runtime.Caller(0)
	testFilesDir      = path.Join(path.Dir(currFile), "..", "tests")

	gimpPaletteCategoriesTestDir = path.Join(testFilesDir, "gimp-palette-categories")
	gimpPalettesTestDir          = path.Join(testFilesDir, "gimp-palettes")
)

func readFile(t *testing.T, filePath string) string {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		t.Errorf("Invalid file path: %s", filePath)
		return ""
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Failed to read file: %s", filePath)
		return ""
	}
	return string(data)
}

func TestGIMPPaletteCategory_String(t *testing.T) {
	type TestData struct {
		category           GIMPPaletteCategory
		expectedOutputFile string
	}

	getTestFilePath := func(filename string) string {
		return path.Join(gimpPaletteCategoriesTestDir, filename)
	}

	testData := []TestData{
		{ // empty category
			category:           GIMPPaletteCategory{},
			expectedOutputFile: "empty-category.txt",
		},
		{ // category with only name, no colors
			category:           GIMPPaletteCategory{Name: "No colors"},
			expectedOutputFile: "category-with-name-and-no-colors.txt",
		},
		{ // category with no name, and colors
			category: GIMPPaletteCategory{
				Colors: []*Color{
					{
						Name: "Black",
						RGB:  RGB{0, 0, 0},
					},
					{
						Name: "White",
						RGB:  RGB{255, 255, 255},
					},
				},
			},
			expectedOutputFile: "category-with-no-name-and-colors.txt",
		},
		{ // category with name and colors
			category: GIMPPaletteCategory{
				Name: "Basic colors",
				Colors: []*Color{
					{
						Name: "Black",
						RGB:  RGB{0, 0, 0},
					},
					{
						Name: "White",
						RGB:  RGB{255, 255, 255},
					},
					{
						Name: "Red",
						RGB:  RGB{255, 0, 0},
					},
					{
						Name: "Green",
						RGB:  RGB{0, 255, 0},
					},
					{
						Name: "Blue",
						RGB:  RGB{0, 0, 255},
					},
				},
			},
			expectedOutputFile: "category-with-name-and-colors.txt",
		},
	}

	for _, td := range testData {
		expectedOutput := readFile(t, getTestFilePath(td.expectedOutputFile))
		if actual := td.category.String(); actual != expectedOutput {
			t.Errorf("Expected \"%s\", got \"%s\"", expectedOutput, actual)
		}
	}
}

func TestGIMPPalette_String(t *testing.T) {
	type TestData struct {
		palette            GIMPPalette
		expectedOutputFile string
	}

	getTestFilePath := func(filename string) string {
		return path.Join(gimpPalettesTestDir, filename)
	}

	testData := []TestData{
		{ // empty palette
			palette:            GIMPPalette{Name: "Empty Palette"},
			expectedOutputFile: "empty-palette.gpl",
		},
		{ // palette with one category, no category name
			palette: GIMPPalette{
				Name: "One category, no category name",
				Categories: []*GIMPPaletteCategory{
					{
						Colors: []*Color{
							{
								Name: "Black",
								RGB:  RGB{0, 0, 0},
							},
							{
								Name: "White",
								RGB:  RGB{255, 255, 255},
							},
						},
					},
				},
			},
			expectedOutputFile: "one-category-no-category-name.gpl",
		},
		{ // palette with two categories, no category names
			palette: GIMPPalette{
				Name: "Two categories, no category names",
				Categories: []*GIMPPaletteCategory{
					{
						Colors: []*Color{
							{
								Name: "Black",
								RGB:  RGB{0, 0, 0},
							},
							{
								Name: "White",
								RGB:  RGB{255, 255, 255},
							},
						},
					},
					{
						Colors: []*Color{
							{
								Name: "Red",
								RGB:  RGB{255, 0, 0},
							},
							{
								Name: "Green",
								RGB:  RGB{0, 255, 0},
							},
							{
								Name: "Blue",
								RGB:  RGB{0, 0, 255},
							},
						},
					},
				},
			},
			expectedOutputFile: "two-categories-no-category-names.gpl",
		},
		{ // palette with two categories, with category names
			palette: GIMPPalette{
				Name: "Two categories, with category names",
				Categories: []*GIMPPaletteCategory{
					{
						Name: "Black and White",
						Colors: []*Color{
							{
								Name: "Black",
								RGB:  RGB{0, 0, 0},
							},
							{
								Name: "White",
								RGB:  RGB{255, 255, 255},
							},
						},
					},
					{
						Name: "Red, Green, and Blue",
						Colors: []*Color{
							{
								Name: "Red",
								RGB:  RGB{255, 0, 0},
							},
							{
								Name: "Green",
								RGB:  RGB{0, 255, 0},
							},
							{
								Name: "Blue",
								RGB:  RGB{0, 0, 255},
							},
						},
					},
				},
			},
			expectedOutputFile: "two-categories-with-category-names.gpl",
		},
	}

	for _, td := range testData {
		expectedOutput := readFile(t, getTestFilePath(td.expectedOutputFile))
		if actual := td.palette.String(); actual != expectedOutput {
			t.Errorf("Expected \"%s\", got \"%s\"", expectedOutput, actual)
		}
	}
}
