package main

import (
	"path"
	"runtime"
)

var (
	_, file, _, _ = runtime.Caller(0)
	dataDir       = path.Join(path.Dir(file), "..", "data")
)

var (
	MTRLogoYAMLFile      = path.Join(dataDir, "mtr-logo.yaml")
	MTRSystemMapYAMLFile = path.Join(dataDir, "mtr-system-map.yaml")
)
