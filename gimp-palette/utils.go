package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadAndUnmarshalYAML[T interface{}](filePath string) (*T, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data *T
	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
