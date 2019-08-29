package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	pkg string = `package %s

import ()`
	pkgTest string = `package %s

import (
    "testing"
)
`
	execText string = `package main

import ()

func main() {
    // your code here
}`
	execTest string = `package main

import (
    "testing"
)
`
	ignore string = ".env\n.DS_Store\n"
	readme string = "## %s"
)

var logger *Log

type project struct {
	appType string
	appName string
	saveDir string
}

func init() {
	logger = StartLog()
}

func main() {
	proj, err := newProject().
		createDirs().
		createProject()

	if err != nil {
		logger.E("Error creating project: %s", err.Error())
		os.Exit(1)
	}

	writeCommonFiles(proj)
}

func newProject() *project {
	if len(os.Args) < 3 {
		logger.E("Too few arguments: %v provided, need 3", len(os.Args))
		os.Exit(1)
	}
	if len(os.Args) > 3 {
		logger.E("Too many arguments: %v provided, only 3 required", len(os.Args))
		os.Exit(1)
	}

	// appType: exec or package, i.e., gonew exec / gonew package
	appType := os.Args[1]
	// appName: name of the application, i.e., gonew exec scraper
	appName := os.Args[2]

	// optional save location, i.e., gonew package httpclient --save ~/develop
	saveDir := flag.String("save", ".", "Directory to save new application to")
	flag.Parse()

	return &project{
		appType: appType,
		appName: appName,
		saveDir: filepath.Join(must(filepath.Abs(*saveDir)), appName),
	}
}

func (proj *project) createDirs() *project {
	err := os.MkdirAll(proj.saveDir, os.ModePerm)
	if err != nil {
		logger.E("Error creating directories: %s", err.Error())
		os.Exit(1)
	}

	logger.I("Created %s", proj.saveDir)
	return proj
}

func (proj *project) createProject() (*project, error) {
	var err error
	if proj.appType == "exec" {
		err = proj.createExec()
		return proj, err
	}

	if proj.appType == "package" {
		err = proj.createPackage()
		return proj, err
	}

	logger.E("Error writing application: %s", err.Error())
	return proj, err
}

func (proj *project) createExec() error {
	f, err := os.Create(filepath.Join(proj.saveDir, "main.go"))
	if err != nil {
		return err
	}
	defer closeFile(f)()

	_, err = f.Write([]byte(execText))
	if err != nil {
		return err
	}

	return nil
}

func (proj *project) createPackage() error {
	return ioutil.WriteFile(
		filepath.Join(proj.saveDir, fmt.Sprintf("%s.go", proj.appName)),
		[]byte(fmt.Sprintf(pkg, proj.appName)),
		0644)

	f, err := os.Create(filepath.Join(proj.saveDir, fmt.Sprintf("%s.go", proj.appName)))
	if err != nil {
		return err
	}
	defer closeFile(f)()

	_, err = f.Write([]byte(fmt.Sprintf(pkg, strings.ToLower(proj.appName))))
	if err != nil {
		return err
	}

	return nil

}

func writeCommonFiles(proj *project) {
	err := ioutil.WriteFile(filepath.Join(proj.saveDir, ".gitignore"), []byte(ignore), os.ModePerm)
	if err != nil {
		logger.E("Error writing .gitignore: %s", err.Error())
		os.Exit(1)
	}

	err = ioutil.WriteFile(filepath.Join(proj.saveDir, "README.md"), []byte(fmt.Sprintf(readme, proj.appName)), os.ModePerm)
	if err != nil {
		logger.E("Error writing README.md: %s", err.Error())
		os.Exit(1)
	}
}

func must(str string, err error) string {
	if err != nil {
		logger.E("Error: %s", err.Error())
		os.Exit(1)
	}

	return str
}

func closeFile(f *os.File) func() {
	return func() {
		err := f.Close()
		if err != nil {
			logger.E("Error closing file: %s", err.Error())
		}
	}
}
