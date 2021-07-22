package gen

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Generator interface {
	Generate() error
}

type GoGenerator struct {
	module string
}

func NewGoGenerator(mod string) *GoGenerator {
	return &GoGenerator{module: mod}
}

func (g *GoGenerator) Generate() error {
	// github.com/joeyscat/oook
	moduleName := g.module

	if !validateModuleName(moduleName) {
		return errors.New(fmt.Sprintf("module name [%s] invalid", moduleName))
	}

	projectName := moduleName

	if strings.Contains(moduleName, "/") {
		ss := strings.Split(moduleName, "/")
		projectName = ss[len(ss)-1]
	}
	if !validateProjectName(projectName) {
		return errors.New(fmt.Sprintf("project name [%s] invalid", projectName))
	}

	return initProject(moduleName, projectName)
}

func validateModuleName(module string) bool {
	// TODO
	if strings.TrimSpace(module) == "" {
		return false
	}

	return true
}

func validateProjectName(project string) bool {
	// TODO
	if strings.TrimSpace(project) == "" {
		return false
	}

	return true
}

func initProject(moduleName, projectName string) error {
	err := os.Mkdir(projectName, 0755)
	if err != nil {
		return err
	}

	mainFile := path.Join(projectName, "main.go")
	err = ioutil.WriteFile(mainFile, []byte(mainFileStr), 0755)
	if err != nil {
		return err
	}

	var output []byte

	// go mod init ...
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = path.Join(projectName)

	if output, err = cmd.Output(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, output)
		return err
	}

	// git init
	cmd = exec.Command("git", "init")
	cmd.Dir = path.Join(projectName)

	if output, err = cmd.Output(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, output)
		return err
	}

	return nil
}

var mainFileStr = `
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

`
