package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Based on the answer here: https://stackoverflow.com/a/49057861
func createBackup() error {
	pathToZip := `C:\Program Files\Intel\Intel Arc Control\resource`
	destinationPath := `backup.zip`

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	myZip := zip.NewWriter(destinationFile)
	defer myZip.Close()

	err = filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(pathToZip, filePath)
		if err != nil {
			return err
		}

		zipFile, err := myZip.Create(relPath)
		if err != nil {
			return err
		}

		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer fsFile.Close()

		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}

		fsFile.Close()
		return nil
	})
	myZip.Close()

	return err
}

func restoreBackup() error {
	reader, err := zip.OpenReader(`backup.zip`)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		filePath := filepath.Join(`C:\Program Files\Intel\Intel Arc Control\resource`, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}

		zipFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zipFile.Close()

		newFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, zipFile)
		if err != nil {
			return err
		}
	}

	fmt.Println("Backup restored.")
	return nil
}
