package main

import "fmt"

func main() {
	data, err := ReadMTRLogoYAMLFile()
	if err != nil {
		panic(err)
	}

	for k, v := range data {
		fmt.Printf("%s: %s [%d %d %d]\n", k, v.Name, v.RGB[0], v.RGB[1], v.RGB[2])
	}
}
