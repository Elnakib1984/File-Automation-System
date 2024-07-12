package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "C:\\" // Change this to the root directory you want to start scanning from (e.g., "C:\\" for Windows)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Log the error and continue walking
			log.Printf("error accessing path %s: %v\n", path, err)
			return nil // skip the directory or file causing the error
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".pdf") {
			fmt.Println("Found PDF:", path) // print the file path if it is a PDF
		}
		return nil // continue walking
	})

	if err != nil {
		log.Fatalf("error walking the path %v: %v\n", root, err)
	}
}
