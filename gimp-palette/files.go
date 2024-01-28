package main

import (
	"os"
	"path"
	"runtime"

	"github.com/gbarail/mtr-colors/types"
	"gopkg.in/yaml.v3"
)

var (
	_, file, _, _ = runtime.Caller(0)
	dataDir       = path.Join(path.Dir(file), "..", "data")
)

var (
	MTRLogoYAMLFile      = path.Join(dataDir, "mtr-logo.yaml")
	MTRSystemMapYAMLFile = path.Join(dataDir, "mtr-system-map.yaml")
)

const GIMPPaletteExtension = "gpl"

func ReadMTRLogoYAMLFile() (types.MTRLogoData, error) {
	bytes, err := os.ReadFile(MTRLogoYAMLFile)
	if err != nil {
		return nil, err
	}

	data := types.MTRLogoData{}
	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GenerateMTRLogoCategory(data types.MTRLogoData) types.Category {
	category := types.Category{
		Name:   "MTR logo",
		Colors: []*types.Color{},
	}

	for _, v := range data {
		category.Colors = append(category.Colors, &types.Color{
			Name: v.Name,
			RGB:  v.RGB,
		})
	}

	return category
}
