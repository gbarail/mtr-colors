package main

import (
	"path"
	"runtime"
	"testing"
)

var (
	_, currFile, _, _ = runtime.Caller(0)
	currDir           = path.Dir(currFile)

	yamlFilesDir = path.Join(currDir, "tests", "yaml")
)

func TestReadAndUnmarshalYAML(t *testing.T) {
	// test invalid YAML path
	func() {
		_, err := ReadAndUnmarshalYAML[interface{}]("/")
		if err == nil {
			t.Errorf("Should have returned an error because the YAML file path is invalid")
		}
	}()

	// test invalid YAML file
	func() {
		yamlFile := path.Join(yamlFilesDir, "invalid.yaml")

		_, err := ReadAndUnmarshalYAML[interface{}](yamlFile)
		if err == nil {
			t.Errorf("Should have returned an error because the YAML file is invalid")
		}
	}()

	// test custom data structure
	func() {
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
		if err != nil {
			t.Error(err)
			return
		}
		if *a != expected {
			t.Errorf("expected %v, got %v", expected, *a)
		}
	}()
}
