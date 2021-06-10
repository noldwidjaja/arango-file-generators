package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateUseCase(directory string, name string, hasRepository bool) {
	camelName := snakeCaseToCamelCase(name)
	titleName := strings.Title(camelName)
	path := directory + "usecase/"

	f, err := os.Create(path + name + "_usecase.go")
	check(err)

	defer f.Close()

	var content []byte
	if hasRepository {
		content = []byte(`
package usecase

type ` + titleName + `UseCaseInterface interface {
	
}

type ` + camelName + `UseCase struct {
	` + camelName + `Repository repository.` + titleName + `RepositoryInterface
}

func New` + titleName + `UseCase(
	` + camelName + `Repository repository.` + titleName + `RepositoryInterface,
) ` + titleName + `UseCaseInterface {
	return &` + camelName + `UseCase{
		` + camelName + `Repository:` + camelName + `Repository,
	}
}
		`)
	} else {
		content = []byte(`
package usecase

type ` + titleName + `UseCaseInterface interface {}

type ` + camelName + `UseCase struct {}

func New` + titleName + `UseCase() ` + titleName + `UseCaseInterface {
	return &` + camelName + `UseCase{}
}
		`)
	}

	err = ioutil.WriteFile(path+name+"_usecase.go", content, 0777)
	check(err)
}
