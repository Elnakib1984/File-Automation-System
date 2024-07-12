package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	root := "/" // Change this to the root directory you want to start scanning from

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err) // log the error if any
			return nil       // continue walking
		}
		fmt.Println(path) // print the file/directory path
		return nil        // continue walking
	})

	if err != nil {
		log.Fatalf("error walking the path %v: %v\n", root, err)
	}
}
