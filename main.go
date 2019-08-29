package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	pkg      string = "package %s\n"
	execText string = "package main\n\nfunc main() {\n    // your code here\n}"
	ignore   string = ".env\n.DS_Store\n"
	readme   string = "## %s"
)

type project struct {
	appType string
	appName string
	saveDir string
}

func main() {
	proj, err := newProject().createDirs().createProject()
	if err != nil {
		os.Exit(1)
	}

	writeCommonFiles(proj)
}

func newProject() *project {
	if len(os.Args) < 3 {
		os.Exit(1)
	}

	appType := os.Args[1]
	appName := os.Args[2]

	saveDir := flag.String("save", ".", "Directory to save new application to")
	flag.Parse()

	return &project{
		appType: appType,
		appName: appName,
		saveDir: filepath.Join(must(filepath.Abs(*saveDir)), appName),
	}
}

func (proj *project) createDirs() *project {
	e := os.MkdirAll(proj.saveDir, os.ModeDir)
	if e != nil {
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Created %s", proj.saveDir))
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

	fmt.Println("Could not write application")
	return proj, err
}

func (proj *project) createExec() error {
	err := ioutil.WriteFile(filepath.Join(proj.saveDir, "main.go"), []byte(execText), 0644)

	return err
}

func (proj *project) createPackage() error {
	err := ioutil.WriteFile(filepath.Join(proj.saveDir, "main.go"), []byte(fmt.Sprintf(pkg, proj.appName)), 0644)

	return err
}

func writeCommonFiles(proj *project) {
	err := ioutil.WriteFile(filepath.Join(proj.saveDir, ".gitignore"), []byte(ignore), os.ModePerm)
	if err != nil {
		os.Exit(1)
	}

	err = ioutil.WriteFile(filepath.Join(proj.saveDir, "README.md"), []byte(fmt.Sprintf(readme, proj.appName)), os.ModePerm)
	if err != nil {
		os.Exit(1)
	}
}

func must(str string, err error) string {
	if err != nil {
		os.Exit(1)
	}

	return str
}
