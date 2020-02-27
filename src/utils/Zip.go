package utils

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//Zip create an zip file from src to dst
func Zip(src string, dst string) error {
	// Get a Buffer to Write To
	outFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	err = addFiles(w, src, "")

	if err != nil {
		fmt.Println(err)
		return err
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func addFiles(w *zip.Writer, basePath, baseInZip string) error {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".cha" {
			dat, err := ioutil.ReadFile(basePath + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
				return err
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
				return err
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + "/" + file.Name()

			err := addFiles(w, newBase, baseInZip+file.Name()+"/")
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}
