package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestReadMTRLogoYAMLFile(t *testing.T) {
	data, err := ReadMTRLogoYAMLFile()
	assert.NilError(t, err)

	assert.Equal(t, len(data), 3) // 3 colors in MTR logo
}

func TestReadMTRSystemMapYAMLFile(t *testing.T) {
	data, err := ReadMTRSystemMapYAMLFile()
	assert.NilError(t, err)

	assert.Assert(t, len(data.LinesColors) > 0)
	assert.Assert(t, len(data.MiscellaneousColors) > 0)
}
