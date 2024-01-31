package main

import (
	"path"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGeneratePalette(t *testing.T) {
	palette := GeneratePalette()

	assert.Equal(t, palette.Name, PaletteName)
	assert.Equal(t, len(palette.Categories), 3)

	for _, category := range palette.Categories {
		assert.Assert(t, len(category.Colors) > 0)
	}
}

func TestWritePaletteToFile(t *testing.T) {
	palette := GeneratePalette()

	{ // write to root directory, should fail
		_, err := WritePaletteToFile(&palette, "/")
		assert.Assert(t, err != nil)
	}

	{ // write to temporary file, should succeed
		tempDir := t.TempDir()
		tempFile := path.Join(tempDir, "palette.gpl")
		_, err := WritePaletteToFile(&palette, tempFile)
		assert.Assert(t, err == nil)
	}
}
