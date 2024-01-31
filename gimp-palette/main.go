package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/gbarail/mtr-colors/types"
)

const PaletteName = "Hong Kong MTR colors"

var (
	_, filename, _, _ = runtime.Caller(0)
	outputPath        = path.Join(path.Dir(filename), "output", "mtr-colors.gpl")
)

func GeneratePalette() types.GIMPPalette {
	palette := types.GIMPPalette{
		Name: PaletteName,
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

	return palette
}

// Write palette to file.
// Returns two values:
// * In case of error, return the empty file path, and error.
// * In case of success, return the actual absolute file path, and nil error.
func WritePaletteToFile(palette *types.GIMPPalette, filePath string) (string, error) {
	filePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = file.WriteString(palette.String())
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func main() {
	palette := GeneratePalette()

	outputPath, err := WritePaletteToFile(&palette, outputPath)
	if err != nil {
		panic(err)
	}

	log.Printf("Written palette to \"%s\".\n", outputPath)
}
