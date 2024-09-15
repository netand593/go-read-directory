package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func findYamlFilesByDir(dir string) (map[string][]string, error) {
	yamlFiles := make(map[string][]string)

	// Walk the directory to find .yaml or .yml files
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a file and has a .yaml or .yml extension
		if !d.IsDir() && (filepath.Ext(d.Name()) == ".yaml" || filepath.Ext(d.Name()) == ".yml") {
			// Get the directory name (without the file)
			dirPath, _ := filepath.Rel(dir, filepath.Dir(path))

			// Append the file name to the slice in the map
			yamlFiles[dirPath] = append(yamlFiles[dirPath], d.Name())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return yamlFiles, nil
}

func GenerateCIFiles() {
	dir := "./environments" // Replace with your directory path

	yamlFiles, err := findYamlFilesByDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(yamlFiles) > 0 {
		fmt.Println("YAML files found:")
		for dirPath, files := range yamlFiles {
			fmt.Printf("Directory: %s\n", dirPath)
			for _, file := range files {
				fmt.Printf("  - %s\n", file)
			}
		}
	} else {
		fmt.Println("No YAML files found.")
	}
	for k, v := range yamlFiles {
		fmt.Println(k, v)
	}
}

func main() {
	dir := "./environments" // Replace with your directory path

	yamlFiles, err := findYamlFilesByDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(yamlFiles) > 0 {
		fmt.Println("YAML files found:")
		for dirPath, files := range yamlFiles {
			fmt.Printf("Directory: %s\n", dirPath)
			for _, file := range files {
				fmt.Printf("  - %s\n", file)
			}
		}
	} else {
		fmt.Println("No YAML files found.")
	}
	for k, v := range yamlFiles {
		fmt.Println(k, v)
	}
}
