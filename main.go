package main

import (
	"os"
	"path/filepath"
)

func main() {
	args := parseArgs()
	path := createDir(args.saveDir)
}

func createDir(dir string) string {
	path, _ := filepath.Abs(dir)
	e := os.MkdirAll(path, os.ModeDir)
	if e != nil {
		os.Exit(1)
	}

	return path
}

func writeFiles(app string) {
	if app == "executable" {
		//
	}
}

func createExec() {
	//
}

func createPackage() {
	//
}
