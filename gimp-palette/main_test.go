package main

import (
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
