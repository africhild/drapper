// config function to get file directory from the config object
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Root    string
	FileDir string
}

// Path: config/config.go
// GetFileDir returns the file directory from the config object
func (c *Config) GetFileDir() string {
	// get the file directory from the config object
	fileDir := c.Root + "/" + c.FileDir

	// if the file directory is empty, use the current working directory
	if fileDir == "" {
		fileDir = "."
	}

	// if the file directory is relative, make it absolute
	if !filepath.IsAbs(fileDir) {
		// get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// make the file directory absolute
		fileDir = filepath.Join(cwd, fileDir)
	}

	// if the file directory is a symlink, resolve it
	fileDir, err := filepath.EvalSymlinks(fileDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if the file directory is a directory, return it
	if fi, err := os.Stat(fileDir); err == nil && fi.IsDir() {
		return fileDir
	}

	// if the file directory is a file, return the directory containing it
	if fi, err := os.Stat(fileDir); err == nil && !fi.IsDir() {
		return filepath.Dir(fileDir)
	}

	// if the file directory does not exist, create it
	if err := os.MkdirAll(fileDir, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// return the file directory
	return fileDir
}

// GetFile returns the file  from the config object and the given file category
func (c *Config) GetFile(template string, ext string) string {
	// get the file directory from the config object
	fileDir := c.GetFileDir()

	// get the file name from the config object
	fileName := template + "." + ext

	// if the file name is empty, use the root
	if fileName == "" {
		fileName = c.Root
	}

	// if the file name is relative, make it absolute
	// if !filepath.IsAbs(fileName) {
	// 	// get the current working directory
	// 	cwd, err := os.Getwd()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	// make the file name absolute
	// 	fileName = filepath.Join(cwd, fileName)
	// }

	// // if the file name is a symlink, resolve it
	// fileName, err := filepath.EvalSymlinks(fileName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// // if the file name is a directory, return the file in it
	// if fi, err := os
	return filepath.Join(fileDir, fileName)
}
