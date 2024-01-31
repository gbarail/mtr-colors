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
