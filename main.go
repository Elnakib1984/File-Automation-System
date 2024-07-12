package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var N string = "Novels"
var ph string = "Philosophy"
var H string = "History"
var oop string = "oop"
var DSA string = "Data Structure"

var oc = os.Create

var pl = fmt.Println

func scanPdf() {

}

func main() {

	srcDir := "E:/semester 2/m213" //source
	destDir := "E:/book/oop"       //destination

	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.MkdirAll(destDir, os.ModePerm)
	}

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), "oop.pdf") {
			oldPath := path
			newPath := filepath.Join(destDir, info.Name())
			err := os.Rename(oldPath, newPath)
			if err != nil {
				return err
			}
			fmt.Printf("Moved: %s -> %s\n", oldPath, newPath)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PDF files ending with 'oop' moved successfully!")
}

/*

Some git commands

# Step 1: Clone the main repository
git clone https://github.com/main-repo.git
cd main-repo

# Step 2: Add the fork as a remote
git remote add fork https://github.com/your-forked-repo.git

# Step 3: Fetch the fork
git fetch fork

# Step 4: Check out the master branch
git checkout master

# Step 5: Merge the fork into master
git merge fork/branch-name

# Step 6: Resolve any conflicts
# Open files, make necessary changes, then
git add conflict-file

# Step 7: Commit the merge
git commit

# Step 8: Push the changes
git push origin master
*/
