package main

import (
	"path"
	"runtime"
	"testing"

	"gotest.tools/v3/assert"
)

var (
	_, currFile, _, _ = runtime.Caller(0)
	currDir           = path.Dir(currFile)

	yamlFilesDir = path.Join(currDir, "tests", "yaml")
)

func TestReadAndUnmarshalYAML(t *testing.T) {
	{ // test invalid YAML path
		_, err := ReadAndUnmarshalYAML[interface{}]("/")
		assert.Assert(t, err != nil)
	}

	{ // test invalid YAML file
		yamlFile := path.Join(yamlFilesDir, "invalid.yaml")

		_, err := ReadAndUnmarshalYAML[interface{}](yamlFile)
		assert.Assert(t, err != nil)
	}

	{ // test custom data structure
		yamlFile := path.Join(yamlFilesDir, "a.yaml")

		type A struct {
			A int
			B string
			C [2]int
		}

		expected := A{
			A: 1,
			B: "b",
			C: [2]int{3, 4},
		}

		var a *A
		a, err := ReadAndUnmarshalYAML[A](yamlFile)
		assert.NilError(t, err)
		assert.DeepEqual(t, *a, expected)
	}
}
